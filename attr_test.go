package tactics

import "testing"

func TestNormal(t *testing.T) {
	outputLevel = 0
	c := Hp4000_Ar100()
	c.Simulate("TestNormal")
	res := c.result
	if res.aliveTime != 8 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}

func TestAR(t *testing.T) {
	outputLevel = 0
	c := Hp4000_Ar100()
	c.Add(AR(50))
	c.Simulate("TestAR")
	res := c.result
	if res.aliveTime != 10 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}

func TestHpAmp1(t *testing.T) {
	outputLevel = 0
	c := Hp4000_Ar100()
	c.Add(HpAmp(25))
	c.Simulate("TestHpAmp")
	res := c.result
	if res.aliveTime != 10 {
		t.Errorf("aliveTime = %d, want 10", res.aliveTime)
	}
	if res.postDmg != 5000 {
		t.Errorf("postDmg = %d, want 5000", res.postDmg)
	}
}

func TestHpAmp2(t *testing.T) {
	outputLevel = 0
	c := Hp4000_Ar100()
	c.Add(HpAmp(50))
	c.Simulate("TestHpAmp")
	res := c.result
	if res.aliveTime != 12 {
		t.Errorf("aliveTime = %d, want 12", res.aliveTime)
	}
	if res.postDmg != 6000 {
		t.Errorf("postDmg = %d, want 5000", res.postDmg)
	}
}

func TestDR1(t *testing.T) {
	outputLevel = 0
	c := Hp4000_Ar100()
	c.Add(DR(50))
	c.Simulate("TestDR1")
	res := c.result
	if res.aliveTime != 16 {
		t.Errorf("aliveTime = %d, want 16", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}

func TestDR2(t *testing.T) {
	outputLevel = 0
	c := Hp4000_Ar100()
	c.Add(DR(75))
	c.Simulate("TestDR2")
	res := c.result
	if res.aliveTime != 32 {
		t.Errorf("aliveTime = %d, want 16", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}

func TestDR3(t *testing.T) {
	outputLevel = 0
	c := Hp4000_Ar100()
	c.Add(DR(90))
	c.Simulate("TestDR3")
	res := c.result
	if res.aliveTime != 80 {
		t.Errorf("aliveTime = %d, want 16", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}

func TestBlock1(t *testing.T) {
	outputLevel = 0
	c := Hp4000_Ar100()
	c.Add(BLK(100))
	c.Simulate("TestBlk")
	res := c.result
	if res.aliveTime != 10 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}

func TestBlock2(t *testing.T) {
	outputLevel = 0
	c := Hp4000_Ar100()
	c.Add(BLK(250))
	c.Simulate("TestBlk")
	res := c.result
	if res.aliveTime != 16 {
		t.Errorf("aliveTime = %d, want 8", res.aliveTime)
	}
	if res.postDmg != 4000 {
		t.Errorf("postDmg = %d, want 4000", res.postDmg)
	}
}
