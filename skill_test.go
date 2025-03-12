package tactics

import (
	"testing"
)

func TestHealBase(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Heal(70, 500)
	c.Simulate("TestHealBase")
	res := c.result
	if res.aliveTime != 14 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 7000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestHealAmp(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Heal(70, 400)
	c.Add(HealAmp(25))
	c.Simulate("TestHealAmp")
	res := c.result
	if res.aliveTime != 14 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 7000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestWound(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Heal(70, 745)
	c.Wound = true
	c.Simulate("TestWound")
	res := c.result
	if res.aliveTime != 14 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 7000-6 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestHealMax(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.HealMax(140, 25)
	c.Simulate("TestHealMax")
	res := c.result
	if res.aliveTime != 12 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 6000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestHealLost(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.MixHeal(35*6, 0, 0, 10)
	c.Simulate("TestHealLost")
	res := c.result
	if res.aliveTime != 9 {
		t.Errorf("aliveTime = %d, want 9", res.aliveTime)
	}
	if res.postDmg != 4300 {
		t.Errorf("postDmg = %d, want 4300", res.postDmg)
	}
}

func TestHealLost2(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.MixHeal(35*6, 0, 0, 20)
	c.Simulate("TestHealLost2")
	res := c.result
	if res.aliveTime != 10 {
		t.Errorf("aliveTime = %d, want 9", res.aliveTime)
	}
	if res.postDmg != 4600 {
		t.Errorf("postDmg = %d, want 4300", res.postDmg)
	}
}

func TestHealMix(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.MixHeal(70, 100, 10, 0)
	c.Simulate("TestHealMix")
	res := c.result
	if res.aliveTime != 14 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 7000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestGrowMax0(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.MixGrow(35*6, 0, 25)
	c.Add(HealAmp(10000))
	c.Simulate("TestGrowMax0")
	res := c.result
	if res.aliveTime != 10 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 5000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestGrowMax1(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	c.health = 2000
	c.Add(HpAmp(100))
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.MixGrow(35*6, 0, 25)
	c.Add(HealAmp(10000))
	c.Simulate("TestGrowMax1")
	res := c.result
	if res.aliveTime != 10 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 5000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestGrowHp(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.MixGrow(35*6, 1000, 0)
	c.Add(HealAmp(10000))
	c.Simulate("TestGrowHp")
	res := c.result
	if res.aliveTime != 10 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 5000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestShield(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Shield(35, 500, 4)
	c.Simulate("TestShield")
	res := c.result
	if res.aliveTime != 15 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 4000+500*7 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestShieldAmp(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Shield(35, 400, 4)
	c.Add(HealAmp(25))
	c.Simulate("TestShieldAmp")
	res := c.result
	if res.aliveTime != 15 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 4000+500*7 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestShieldMax(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.MixShield(35, 0, 10, 4)
	c.Add(HealAmp(25))
	c.Simulate("TestShieldMax")
	res := c.result
	if res.aliveTime != 15 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 4000+500*7 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func TestShieldTime(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Shield(70, 9999, 2)
	c.Simulate("TestShieldTime")
	res := c.result
	if res.aliveTime != 14 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 7000 {
		t.Errorf("postDmg = %d, want 8000", res.postDmg)
	}
}

func TestApAmp(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Shield(35, 100, 4)
	c.Add(HealAmp(100))
	c.Add(AP(150))
	c.Simulate("TestApAmpShield")
	res := c.result
	if res.aliveTime != 15 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 4000+500*7 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}

	Level(0)
	c = Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.MixGrow(35*6, 500, 0)
	c.Add(AP(100))
	c.Add(HealAmp(10000))
	c.Simulate("TestApAmpGrowHp")
	res = c.result
	if res.aliveTime != 10 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 5000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}

	Level(0)
	c = Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.Heal(70, 100)
	c.Add(HealAmp(100))
	c.Add(AP(150))
	c.Simulate("TestApAmpHeal")
	res = c.result
	if res.aliveTime != 14 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 7000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}

func Test2Stage(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	// 每秒回蓝35点，普攻10点，挨打回25点
	c.MixGrow(35, 0, 25)
	c.HealMax(70, 10)
	c.Simulate("Test2Stage")
	res := c.result
	if res.aliveTime != 17 {
		t.Errorf("aliveTime = %d, want 14", res.aliveTime)
	}
	if res.postDmg != 4000+500*7+1000 {
		t.Errorf("postDmg = %d, want 7000", res.postDmg)
	}
}
