package test

import (
	o "tactics"
	"testing"
)

func Test13(t *testing.T) {
	o.Level(0)
	// todo修复法力锁错误
	reni := o.Champ(1530, 50, 40).
		MixHeal(100, 325, 15, 0).Swing(2)
	reni.Add(o.HpAmp(80))
	reni.Simulate("2⭐ Reni")

	rell := o.Champ(2592, 45+25*3, 40).
		Shield(90, 400, 4)
	rell.Add(o.ManaAmp(50))
	rell.Simulate("3⭐ Rell")

	amumu := o.Champ(1500, 35+100, 50).
		Watcher()
	amumu.Add(o.BLK(25))
	amumu.Simulate("3⭐ Amumu")

	singed := o.Champ(2106, 40+25*3, 0).
		Buff(50, 4, o.DR(50))
	singed.Simulate("3⭐ Singed")

	illaoi := o.Champ(1980, 60+42*3, 65).
		Grow(125, 4*75).Buff(0, 3, o.DR(50)).Merge()
	illaoi.Simulate("2⭐ Illaoi")

	garen := o.Champ(1800, 60, 60).
		Watcher().
		MixShield(120, 220, 15, 4).
		Garen()
	garen.Simulate("2⭐ Garen")

	scar := o.Champ(2592, 50, 90).
		Watcher().
		Fire().
		MixHeal(150, 330, 0, 0)
	scar.Simulate("3⭐ Scar")

	scar0 := o.Champ(1440, 50, 90).
		Watcher().
		Fire().
		MixHeal(150, 280, 0, 0)
	scar0.Simulate("2⭐ Scar")

	loris := o.Champ(2754, 50+25*3, 50).
		MixShield(100, 900, 0, 4)
	loris.Simulate("3⭐ Loris")

	loris0 := o.Champ(1530, 50+25*3, 50).
		MixShield(100, 700, 0, 4)
	loris0.Simulate("2⭐ Loris")
}
