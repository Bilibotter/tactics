package test

import (
	o "tactics"
	"testing"
)

func Test12(t *testing.T) {
	o.Level(0)
	wukong := o.Champ(1800, 50+20, 20).
		Shield(70, 460, 4).
		Preserver(6, 3)
	wukong.Simulate("2⭐ WuKong")

	wukong0 := o.Champ(3240, 50+20, 20).
		Shield(70, 550, 4).
		Preserver(6, 3)
	wukong0.Simulate("3⭐ WuKong")

	morgana := o.Champ(1980, 60, 40).
		Shield(110, 525, 4).
		Preserver(12, 3)
	morgana.Simulate("2⭐ Morgana")

	rakan := o.Champ(1890, 60, 40).
		MixShield(129, 200, 20, 3).
		Preserver(12, 3)
	rakan.Simulate("2⭐ Rakan")

	taric := o.Champ(1980, 60+25, 50).
		Buff(100, 4, o.DR(60))
	taric.Simulate("2⭐ Taric")

	nasus := o.Champ(1980, 60, 0).
		Grow(80, 600).Heal(80, 300).
		Shapeshifter()
	nasus.Simulate("2⭐ Nasus")
}
