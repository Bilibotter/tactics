package tactics

import "testing"

func TestSwingLock(t *testing.T) {
	Level(3)
	c := Champ(4000, 200, 0).
		MixHeal(100, 400, 10, 0).Swing(2)
	c.health = 4000
	c.healthAmp = 100
	c.armor = 200
	c.Wound = false
	c.Shred = false
	c.Simulate("TestSwingLock")
}
