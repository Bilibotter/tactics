package tactics

import (
	"fmt"
	"sort"
)

var testDmg = 0

var outputLevel = 0 // 3：详细；2：标准；1：重大
var sortOutput = false
var outputs []output
var dmgBefore15, dmgAfter15 = 0, 0 // 15秒前后每秒承受的伤害

type filter_ []Handler

type output struct {
	val       string
	aliveTime int
	postDmg   int
}

type Ground struct {
	*champion_
	DmgRecord  []int // 每秒受到的实际生命值伤害
	AtkTimes   int   // 攻击次数
	CurrenTime int   // 当前时间
	CastTimes  int   // 施法次数
	result     output
}

func SetDmg(dmg1, dmg2 int) {
	dmgBefore15, dmgAfter15 = dmg1, dmg2
}

func Level(level int) {
	outputLevel = level
}

func SortOutput() {
	sortOutput = true
}

func OutputBySort() {
	sort.SliceStable(outputs, func(i, j int) bool {
		if outputs[i].aliveTime == outputs[j].aliveTime {
			return outputs[i].postDmg > outputs[j].postDmg
		}
		return outputs[i].aliveTime > outputs[j].aliveTime
	})
	for _, o := range outputs {
		fmt.Print(o.val)
	}
}

func (g *Ground) filter(event Event) {
	for _, handler := range g.filter_ {
		handler.handle(event, g)
	}
}

func (g *Ground) Simulate(name ...string) *Ground {
	g.buffs = make(map[string]*buff_)
	g.currentHealth = g.healthy()
	if len(name) > 0 {
		g.Name = name[0]
	}
	g.run()
	return g
}

func (g *Ground) run() {
	tick := 0.01
	ticks := 0
	swing := 0
	pre, post := 1, 3 // 折前/折后承伤百分比回蓝
	damage := 500
	if dmgBefore15 > 0 {
		damage = dmgBefore15
	}
	champ := g
	champ.maxHealth = champ.healthy()
	totalDmg := 0
	for g.CurrenTime < 60 || champ.currentHealth > 0 {
		if g.CurrenTime > 15 {
			// 15秒后伤害增加一部分
			damage = 750
			if dmgAfter15 > 0 {
				damage = dmgAfter15
			}
		}
		if testDmg > 0 {
			damage = testDmg
		}
		ticks += 1
		swing += 1
		g.CurrenTime = ticks / 100
		mana := 0 // 本次tick回蓝
		lockingMana := g.lockingMana()
		if float64(swing)*tick*champ.as()-1.0 >= 0 {
			swing = 0
			mana += 10 * champ.manaGain() / 100
			g.AtkTimes += 1
			g.filter(NewE(attackA, g.AtkTimes))
			if outputLevel >= 3 {
				if !lockingMana {
					fmt.Printf("%d秒:第%d次攻击, 回蓝%d点\n", g.CurrenTime, g.AtkTimes, mana)
				} else {
					fmt.Printf("%d秒:第%d次攻击, 法力锁定中\n", g.CurrenTime, g.AtkTimes)
				}
			}
		}
		if ticks%100 == 0 {
			// 先叠加承受伤害叠加的被动再计算折后伤害
			g.filter(NewE(damagedA, g.CurrenTime))
			actual := champ.postDmg(damage)
			if !champ.lose(actual) {
				total := actual
				for _, tmp := range g.DmgRecord {
					total += tmp
				}
				// 当前生命值为负
				total += g.currentHealth
				totalDmg += damage
				val := fmt.Sprintf("%10s承伤总时长%d, 总折后承伤%d, 总承伤%d\n", g.Name, g.CurrenTime, total, totalDmg)
				o := output{val, g.CurrenTime, total}
				if !sortOutput {
					fmt.Print(val)
				} else {
					outputs = append(outputs, o)
				}
				g.result = o
				break
			}
			mana += (damage*pre + actual*post) / 100 * g.manaGain() / 100
			g.DmgRecord = append(g.DmgRecord, actual)
			totalDmg += damage
			showStatus(g)
			showActiveAttach(g)
			// 计算buff持续时间时包含右边界
			g.filter(NewE(timeGoA, g.CurrenTime))
			if outputLevel >= 3 {
				if !g.lockingMana() {
					fmt.Printf("%d秒:折后承伤%d, 回蓝%d点\n", g.CurrenTime, actual, mana)
				} else {
					fmt.Printf("%d秒:折后承伤%d, 法力锁定中\n", g.CurrenTime, actual)
				}
			}
		}
		if !lockingMana {
			champ.currentMana += mana
		}
		// 为避免边界值使得buff收益期超出预期先承伤再施法。比如避免持续2s的buff作用于3次承伤。
		// 为了让技能吃到施法增益，先触发时间性被动再施法。
		if !lockingMana && champ.skill != nil && champ.currentMana >= champ.skill.costMana() {
			if outputLevel >= 2 {
				fmt.Printf("%d秒:第%d次读条施法\n", g.CurrenTime, g.CastTimes+1)
			}
			champ.currentMana -= champ.skill.costMana()
			g.skill.cast()
			g.CastTimes += 1
		}
		g.CurrenTime = ticks / 100
	}
}

func showActiveAttach(ground *Ground) {
	if outputLevel >= 3 {
		num0, num1 := 0, 0
		for _, bf := range ground.buffs {
			if bf.IsValid() {
				num0 += 1
			}
		}
		for _, a := range ground.attach {
			if !a.IsValid() {
				continue
			}
			if _, ok := a.(*passive); ok {
				num1 += 1
			}
		}
		fmt.Printf("%d秒:当前生效buff数%d, 当前生效被动数%d\n", ground.CurrenTime, num0, num1)
	}
}

func showStatus(ground *Ground) {
	if outputLevel < 3 {
		return
	}
	cp := ground.attrs_
	for _, a := range ground.attach {
		if !a.IsValid() {
			continue
		}
		cp.Add(a.attr())
	}
	fmt.Printf("%d秒:生命值:%d, 减伤:%d%%, 双抗:%d, 最大生命值加成:%d\n", ground.CurrenTime, cp.health, 100-cp.dmgTaken, cp.armor, cp.healthAmp-100)
}
