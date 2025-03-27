package tactics

import "fmt"

type techItem int

const (
	RED techItem = iota
	BLUE
	GREEN
	YELLOW
	VIOLET
)

func (g *Ground) Exotech(item techItem, actives ...int) *Ground {
	active := 3
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 3:

	case 5:

	case 7:

	default:
		panic("wrong origin num")
	}
	return g
}

func (g *Ground) Bruiser(actives ...int) *Ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 2:
		g.Add(HpAmp(20))
	case 4:
		g.Add(HpAmp(45))
	case 6:
		g.Add(HpAmp(75))
	default:
		panic("wrong origin num")
	}
	return g
}

// 堡垒卫士
func (g *Ground) Bastion(actives ...int) *Ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	increase := 0
	switch active {
	case 2:
		increase = 18
	case 4:
		increase = 50
	case 6:
		increase = 90
	default:
		panic("wrong origin num")
	}
	g.Add(AR(increase))
	p := &passive{
		trigger: timeLineA,
		start:   11,
		end:     100,
	}
	p.Add(AR(increase))
	g.addPassive(p)
	return g
}

// 重装的盾回蓝
func (g *Ground) Vanguard(actives ...int) *Ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	p := &passive{
		trigger: shieldedA,
	}
	increase := 0
	switch active {
	case 2:
		increase = 16
		p.Add(DR(10))
	case 4:
		increase = 32
		p.Add(DR(10))
	case 6:
		increase = 40
		p.Add(DR(20))
	default:
		panic("wrong origin num")
	}
	p0 := &passive{
		trigger: timeGoA,
	}
	p1 := &passive{
		trigger: timeGoA,
	}
	p0.call = func(ground *Ground) {
		if ground.currentHealth > 0 && ground.currentHealth*100/ground.healthy() < 50 {
			foo := ground.healthy() * increase / 100
			addHealth(foo, ground.champion_)
			p0.once = 2
			if outputLevel >= 3 {
				fmt.Printf("%d秒:第一次触发重装被动\n", ground.CurrenTime)
			}
		}
	}
	p1.call = func(ground *Ground) {
		if ground.currentHealth > 0 && ground.currentHealth*100/ground.healthy() < 50 {
			foo := ground.healthy() * increase / 100
			addHealth(foo, ground.champion_)
			p1.once = 2
			if outputLevel >= 3 {
				fmt.Printf("%d秒:第二次触发重装被动\n", ground.CurrenTime)
			}
		}
	}
	g.addPassive(p)
	g.addPassive(p0)
	g.addPassive(p1)
	return g
}

// 重装的盾回蓝
func (g *Ground) Vanguard0(actives ...int) *Ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	p := &passive{
		trigger: shieldedA,
	}
	increase := 0
	switch active {
	case 2:
		increase = 16
		p.Add(DR(10))
	case 4:
		increase = 32
		p.Add(DR(10))
	case 6:
		increase = 40
		p.Add(DR(20))
	default:
		panic("wrong origin num")
	}
	p0 := &passive{
		trigger: timeGoA,
	}
	p1 := &passive{
		trigger: timeGoA,
	}
	p0.call = func(ground *Ground) {
		shieldMax(increase, 10, true)(ground)
		p0.once = 2
		if outputLevel >= 3 {
			fmt.Printf("%d秒:第一次触发重装被动\n", ground.CurrenTime)
		}
	}
	p1.call = func(ground *Ground) {
		if ground.currentHealth > 0 && ground.currentHealth*100/ground.healthy() < 50 {
			shieldMax(increase, 10, true)(ground)
			p1.once = 2
			if outputLevel >= 3 {
				fmt.Printf("%d秒:第二次触发重装被动\n", ground.CurrenTime)
			}
		}
	}
	g.addPassive(p)
	g.addPassive(p0)
	g.addPassive(p1)
	return g
}

// 战略分析
func (g *Ground) Strategist(actives ...int) *Ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	fac := 3
	increase := 0
	switch active {
	case 2:
		increase = 4
	case 3:
		increase = 6
	case 4:
		increase = 9
	case 5:
		increase = 12
	default:
		panic("wrong origin num")
	}
	g.Add(DR(increase * fac))

	return g
}

func (g *Ground) Syndicate(actives ...int) *Ground {
	active := 3
	if len(actives) != 0 {
		active = actives[0]
	}
	increase := 0
	switch active {
	case 3:
		increase = 100
	case 5:
		increase = 450
	case 7:
		increase = 600
	default:
		panic("wrong origin num")
	}
	g.health += increase
	return g
}

func (g *Ground) Animal(actives ...int) *Ground {
	active := 3
	if len(actives) != 0 {
		active = actives[0]
	}
	increase := 0
	switch active {
	case 3:
		increase = 10
	case 5:
		increase = 20
	case 7:
		increase = 30
	default:
		panic("wrong origin num")
	}
	g.Add(AR(increase))
	return g
}
