package tactics

import "testing"

func Hp4000_Ar100() *Ground {
	c := Champ(4000, 100, 0)
	c.basicSpeed = 1.0
	c.health = 4000
	c.healthAmp = 100
	c.armor = 100
	c.Wound = false
	c.Shred = false
	testDmg = 1000
	return c
}

func TestDRBuff(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	c.Buff(70, 300, DR(50))
	c.Simulate("TestDRBuff")
	res := c.result
	if res.aliveTime != 8+6 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}

func TestARBuff(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	c.Buff(70, 300, AR(200))
	c.Simulate("TestARBuff")
	res := c.result
	if res.aliveTime != 8+6 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}

func TestSwingLock(t *testing.T) {
	Level(0)
	c := Hp4000_Ar100()
	c.MixHeal(10, 0, 0, 0).Swing(30)
	c.Simulate("TestSwingLock")
	res := c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
	if c.CastTimes != 1 {
		t.Errorf("castTimes = %d, want 1", c.CastTimes)
	}
}
