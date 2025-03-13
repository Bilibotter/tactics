package tactics

import "testing"

/*  trigger:
attackA Action
damagedA
beforeCastA
afterCastA
timeGoA
timeLineA // 无事件，标记为只在特定时间段生效
shieldedA
unShieldedA
healedA
healthPercentA
*/

func TestAttackTrigger(t *testing.T) {
	count := 0
	outputLevel = 0
	c := Hp4000_Ar100()
	p := &passive{}
	p.trigger = attackA
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Simulate("TestAttackTrigger")
	res := c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if count != 8 {
		t.Errorf("count = %d, want 8", count)
	}
}

func TestAttackTriggerFreq(t *testing.T) {
	count := 0
	Level(0)
	c := Hp4000_Ar100()
	p := &passive{}
	p.trigger = attackA
	p.call = func(ground *Ground) {
		count++
	}
	p.freq = 2
	c.addPassive(p)
	c.Simulate("TestAttackTriggerFreq")
	res := c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if count != 4 {
		t.Errorf("count = %d, want 4", count)
	}
}

func TestDamagedTrigger(t *testing.T) {
	count := 0
	Level(0)
	c := Hp4000_Ar100()
	p := &passive{}
	p.trigger = damagedA
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Simulate("TestDamagedTrigger")
	res := c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if count != 8 {
		t.Errorf("count = %d, want 8", count)
	}

	count = 0
	Level(0)
	c = Hp4000_Ar100()
	p = &passive{}
	p.trigger = attackA
	p.call = func(ground *Ground) {
		count++
	}
	p.freq = 2
	c.addPassive(p)
	c.Simulate("TestDamagedTriggerFreq")
	res = c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if count != 4 {
		t.Errorf("count = %d, want 4", count)
	}
}

func TestCastTrigger(t *testing.T) {
	count := 0
	Level(0)
	c := Hp4000_Ar100()
	p := &passive{}
	p.trigger = beforeCastA
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Heal(70, 0)
	c.Simulate("TestCastTrigger")
	res := c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if count != 3 {
		t.Errorf("count = %d, want 8", count)
	}

	count = 0
	Level(0)
	c = Hp4000_Ar100()
	p = &passive{}
	p.trigger = beforeCastA
	p.call = func(ground *Ground) {
		count++
	}
	p.freq = 2
	c.addPassive(p)
	c.Heal(35, 0)
	c.Simulate("TestDamagedTriggerFreq")
	res = c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if count != 3 {
		t.Errorf("count = %d, want 4", count)
	}
}

func TestTimeGoTrigger(t *testing.T) {
	count := 0
	Level(0)
	c := Hp4000_Ar100()
	p := &passive{}
	p.trigger = timeGoA
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Simulate("TestDamagedTrigger")
	res := c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if count != 7 {
		t.Errorf("count = %d, want 7", count)
	}

	count = 0
	Level(0)
	c = Hp4000_Ar100()
	p = &passive{}
	p.trigger = timeGoA
	p.call = func(ground *Ground) {
		count++
	}
	p.freq = 2
	c.addPassive(p)
	c.Simulate("TestDamagedTriggerFreq")
	res = c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if count != 3 {
		t.Errorf("count = %d, want 4", count)
	}
}

func TestShieldTrigger(t *testing.T) {
	count := 0
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Shield(35, 500, 4)
	p := &passive{}
	p.trigger = shieldedA
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Simulate("TestShield")
	res := c.result
	if res.aliveTime != 15 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 4000+500*7 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
	if c.CastTimes != 7 {
		t.Errorf("CastTimes = %d, want 7", c.CastTimes)
	}
	if count != 7 {
		t.Errorf("count = %d, want 7", count)
	}

	count = 0
	Level(0)
	c = Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Shield(35, 500, 4)
	p = &passive{}
	p.trigger = shieldedA
	p.freq = 2
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Simulate("TestShield")
	res = c.result
	if res.aliveTime != 15 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 4000+500*7 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
	if c.CastTimes != 7 {
		t.Errorf("CastTimes = %d, want 7", c.CastTimes)
	}
	if count != 3 {
		t.Errorf("count = %d, want 3", count)
	}
}

func TestShieldTrigger0(t *testing.T) {
	count := 0
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Heal(35, 0)
	p := &passive{}
	p.trigger = shieldedA
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Simulate("TestShield")
	res := c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if c.CastTimes != 7 {
		t.Errorf("CastTimes = %d, want 7", c.CastTimes)
	}
	if count != 0 {
		t.Errorf("count = %d, want 0", count)
	}
}

func TestHealTrigger(t *testing.T) {
	Level(0)
	count := 0
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Heal(70, 500)
	p := &passive{}
	p.trigger = healedA
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Simulate("TestHealBase")
	res := c.result
	if res.aliveTime != 14 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 7000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
	if count != 6 {
		t.Errorf("count = %d, want 6", count)
	}

	Level(0)
	count = 0
	c = Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Heal(70, 500)
	p = &passive{}
	p.trigger = healedA
	p.freq = 2
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Simulate("TestHealBase")
	res = c.result
	if res.aliveTime != 14 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 7000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
	if count != 3 {
		t.Errorf("count = %d, want 3", count)
	}
}

func TestHealTrigger0(t *testing.T) {
	count := 0
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Shield(35, 500, 4)
	p := &passive{}
	p.trigger = healedA
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Simulate("TestShield")
	res := c.result
	if res.aliveTime != 15 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 4000+500*7 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
	if c.CastTimes != 7 {
		t.Errorf("CastTimes = %d, want 7", c.CastTimes)
	}
	if count != 0 {
		t.Errorf("count = %d, want 0", count)
	}

	count = 0
	Level(0)
	c = Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Shield(35, 500, 4)
	p = &passive{}
	p.trigger = healedA
	p.freq = 2
	p.call = func(ground *Ground) {
		count++
	}
	c.addPassive(p)
	c.Simulate("TestShield")
	res = c.result
	if res.aliveTime != 15 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 4000+500*7 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
	if c.CastTimes != 7 {
		t.Errorf("CastTimes = %d, want 7", c.CastTimes)
	}
	if count != 0 {
		t.Errorf("count = %d, want 0", count)
	}
}

func TestHealthPercentTrigger(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	p := &passive{}
	p.trigger = healthPercentA
	p.start = 60
	p.end = 100
	p.Add(DR(50))
	c.addPassive(p)
	c.Simulate("TestHealthPercentTrigger")
	res := c.result
	if res.aliveTime != 12 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}

	Level(0)
	c = Hp4000_Ar100()
	p = &passive{}
	p.trigger = healthPercentA
	p.start = 0
	p.end = 50
	p.Add(DR(50))
	c.addPassive(p)
	c.Simulate("TestHealthPercentTrigger")
	res = c.result
	if res.aliveTime != 12 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}

func TestTimeLineTrigger(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	p := &passive{}
	p.trigger = timeLineA
	p.start = 5
	p.end = 16
	p.Add(DR(50))
	c.addPassive(p)
	c.Simulate("TestTimeLineTrigger")
	res := c.result
	if res.aliveTime != 12 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}

	Level(0)
	c = Hp4000_Ar100()
	p = &passive{}
	p.trigger = timeLineA
	p.start = 0
	p.end = 8
	p.Add(DR(50))
	c.addPassive(p)
	c.Simulate("TestHealthPercentTrigger")
	res = c.result
	if res.aliveTime != 12 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}

func TestTriggerOnce(t *testing.T) {
	count := 0
	Level(0)
	c := Hp4000_Ar100()
	p := &passive{}
	p.trigger = attackA
	p.call = func(ground *Ground) {
		count++
	}
	p.once = 1
	c.addPassive(p)
	c.Simulate("TestAttackTrigger")
	res := c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if count != 1 {
		t.Errorf("count = %d, want 1", count)
	}

	count = 0
	Level(3)
	c = Hp4000_Ar100()
	p = &passive{}
	p.trigger = beforeCastA
	p.call = func(ground *Ground) {
		if ground.CastTimes != 2 {
			t.Errorf("CastTimes = %d, want 2", ground.CastTimes)
		}
		count++
	}
	p.freq = 3
	p.once = 1
	c.addPassive(p)
	c.Heal(35, 0)
	c.Simulate("TestAttackTrigger")
	res = c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if count != 1 {
		t.Errorf("count = %d, want 1", count)
	}
}

func TestStackPassive(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	p := c.stackOrigin(AR(1))
	p.trigger = attackA
	p.maxStack = 5
	c.Simulate("TestStackPassive")
	res := c.result
	if res.aliveTime != 9 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	ar := c.armor
	for _, attach := range c.attach {
		if attach.IsValid() {
			ar += attach.attr().armor
		}
	}
	if ar != 105 {
		t.Errorf("armor = %d, want 105", ar)
	}
}

func TestStackPassive0(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	p := c.stackOrigin(AR(1))
	p.trigger = attackA
	p.maxStack = 99
	p.freq = 2
	c.Simulate("TestStackPassive")
	res := c.result
	if res.aliveTime != 9 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	ar := c.armor
	for _, attach := range c.attach {
		if attach.IsValid() {
			ar += attach.attr().armor
		}
	}
	if ar != 104 {
		t.Errorf("armor = %d, want 104", ar)
	}
}
