package tactics

import "fmt"

type champion_ struct {
	attrs_
	filter_
	Name          string
	basicSpeed    float64 // 基础攻速
	currentMana   int     // 当前法力值
	currentHealth int     // 当前生命值
	maxHealth     int     // 最大生命值
	buffs         map[string]*buff_
	locks         []*manaLock
	shields       []*shield_
	Wound         bool // 重伤
	Shred         bool // 破甲
	attach        []attach_
	skill         *skill_
}

func (g *champion_) lockingMana() bool {
	for _, lock := range g.locks {
		if lock.IsValid() {
			return true
		}
	}
	return false
}

func (g *champion_) castRate() int {
	ability, amp := g.ability, g.castAmp
	for _, a := range g.attach {
		if a.IsValid() {
			ability += a.attr().ability
			amp += a.attr().castAmp
		}
	}
	return ability * amp / 100
}

func (g *champion_) manaGain() int {
	amp := g.manaAmp
	for _, a := range g.attach {
		if a.IsValid() {
			amp += a.attr().manaAmp
		}
	}
	return amp
}

func (g *champion_) mitigate() float64 {
	armor, dmgTaken, armorAmp := g.armor, g.dmgTaken, g.armorAmp
	for _, a := range g.attach {
		if a.IsValid() {
			armor += a.attr().armor
			if a.attr().dmgTaken > 0 {
				dmgTaken = dmgTaken * a.attr().dmgTaken / 100
			}
		}
	}
	armor = armor * armorAmp / 100
	if g.Shred {
		armor = armor * 70 / 100
	}
	return float64(dmgTaken) / (100 + float64(armor))
}

func (g *champion_) blocks() int {
	block := g.block
	for _, a := range g.attach {
		if a.IsValid() {
			block += a.attr().block
		}
	}
	return block
}

func (g *champion_) atkSpeed() float64 {
	haste := g.speed
	for _, a := range g.attach {
		if a.IsValid() {
			haste += a.attr().speed
		}
	}
	return g.basicSpeed * float64(haste) / 100
}

func (g *champion_) healthy() int {
	health, amp := g.health, g.healthAmp
	for _, a := range g.attach {
		if !a.IsValid() {
			continue
		}
		g.health += a.attr().health
		g.healthAmp += a.attr().healthAmp
	}
	return health * amp / 100
}

func (g *champion_) shieldAmp() int {
	amp := g.recoverAmp
	for _, a := range g.attach {
		if a.IsValid() {
			amp += a.attr().recoverAmp
		}
	}
	return amp
}

func (g *champion_) healAmp() int {
	amp := g.shieldAmp()
	if g.Wound {
		amp = amp * 67 / 100
	}
	return amp
}

func (g *champion_) as() float64 {
	speed := 100
	for _, a := range g.attach {
		if a.IsValid() {
			speed += a.attr().speed
		}
	}
	return g.basicSpeed * float64(speed) / 100
}

// 计算折后伤害
func (g *champion_) postDmg(preDmg int) int {
	m := g.mitigate()
	dmg := int(float64(preDmg)*m) - g.blocks()
	if dmg <= 0 {
		fmt.Println(preDmg, g.mitigate(), g.blocks())
		panic("damage less than 0")
	}
	return dmg
}

// 实际扣除生命值和护盾
func (g *Ground) lose(dmg int) bool {
	old := dmg
	for _, sh := range g.shields {
		dmg = sh.Taken(dmg)
		if dmg == 0 {
			if outputLevel >= 3 {
				fmt.Printf("护盾完全抵挡%d点伤害, 剩余护盾%d\n", old, g.shieldHealth())
			}
			return true
		}
	}
	g.currentHealth -= dmg
	if outputLevel >= 3 {
		fmt.Printf("%d秒扣除生命值%d, 当前生命值%d, 最大生命值%d\n", g.CurrenTime, dmg, g.currentHealth, g.maxHealth)
	}
	return g.currentHealth > 0
}

func (g *champion_) shieldHealth() int {
	total := 0
	for _, sh := range g.shields {
		if sh.IsValid() {
			total += sh.health
		}
	}
	return total
}
