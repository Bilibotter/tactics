package tactics

import "fmt"

func (g *Ground) attrOrigin(attrs ...*attrs_) *passive {
	p := &passive{}
	for _, attr := range attrs {
		p.Add(attr)
	}
	g.addPassive(p)
	return p
}

func (g *Ground) stackOrigin(attrs ...*attrs_) *passive {
	p := &passive{
		stack: &attrs_{},
	}
	for _, attr := range attrs {
		p.stack.Add(attr)
	}
	g.addPassive(p)
	return p
}

func (g *Ground) buffOrigin(trigger Action, duration int, attrs ...*attrs_) *passive {
	p := &passive{
		trigger: trigger,
	}
	p.call = addBuff(buff(duration, attrs...))
	g.addPassive(p)
	return p
}

// 野火帮
func (g *Ground) Fire() *Ground {
	p := &passive{}
	p.trigger = timeGoA
	p.freq = 6
	p.call = func(ground *Ground) {
		champion := ground.champion_
		lost := 0
		for i := 1; i <= 6; i++ {
			lost += ground.DmgRecord[len(ground.DmgRecord)-i]
		}
		heal := lost * champion.healAmp() / 100 * 20 / 100
		if outputLevel >= 3 {
			fmt.Printf("%d秒:野火恢复%d, 总损失%d\n", ground.CurrenTime, heal, lost)
		}
		addHealth(heal, champion)
	}
	return g.addPassive(p)
}

func (g *Ground) Invoker() *Ground {
	p := &passive{}
	p.trigger = timeGoA
	p.freq = 3
	p.call = func(ground *Ground) {
		champion := g.champion_
		champion.currentMana += 5
	}
	return g.addPassive(p)
}

func (g *Ground) Shurima() *Ground {
	p0 := &passive{
		trigger: timeGoA,
		freq:    8,
		call:    growMax(30),
		once:    1,
	}

	p1 := &passive{
		trigger: timeGoA,
		freq:    4,
		call:    healMax(5),
	}
	return g.addPassive(p0, p1)
}

func (g *Ground) Warden() *Ground {
	p0 := &passive{
		trigger: timeLineA,
		start:   0,
		end:     11,
	}
	p0.dmgTaken = 72

	p1 := &passive{
		trigger: timeLineA,
		start:   11,
		end:     100,
	}
	p1.dmgTaken = 90
	return g.addPassive(p0, p1)
}

func (g *Ground) Watcher() *Ground {
	p0 := &passive{
		trigger: healthPercentA,
		start:   50,
		end:     101,
	}
	p0.dmgTaken = 70

	p1 := &passive{
		trigger: healthPercentA,
		start:   0,
		end:     50,
	}
	p1.dmgTaken = 85
	return g.addPassive(p0, p1)
}

func (g *Ground) Garen() *Ground {
	p := &passive{
		trigger: attackA,
	}
	p.call = healMax(1)
	return g.addPassive(p)
}

func (g *Ground) Annie() *Ground {
	p := &passive{
		trigger: attackA,
	}
	p.call = healMax(5)
	return g.addPassive(p)
}

func (g *Ground) Porcelain() *Ground {
	p := &passive{
		trigger: beforeCastA,
	}
	p.call = addBuff(buff(4, DR(20)))
	return g.addPassive(p)
}

func (g *Ground) Shapeshifter() *Ground {
	g.healthAmp += 30
	return g
}

// 日炎灼烧完全抵消龙牙的治疗，因此不计算
func (g *Ground) Preserver(num, freq int) *Ground {
	p := &passive{}
	p.trigger = timeGoA
	p.freq = freq
	p.call = healMax(num)
	return g.addPassive(p)
}
