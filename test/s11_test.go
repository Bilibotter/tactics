package test

import (
	o "tactics"
	"testing"
)

func Test11(t *testing.T) {
	o.Level(0)

	//illaoi_ := o.Champ(1530, 50, 20).
	//	Shield(110, 400, 4).Heal(0, 60*4).
	//	Swing(4).Merge().
	//	Warden().
	//	Preserver(2, 5)
	//illaoi_.Add(o.AP(20))
	//illaoi_.Simulate("2⭐ Illaoi<!>")
	//
	//illaoi0_ := o.Champ(2754, 50, 20).
	//	Shield(110, 400, 4).Heal(0, 70*4).
	//	Swing(4).Merge().
	//	Warden().Preserver(2, 5)
	//illaoi0_.Add(o.AP(20))
	//illaoi0_.Simulate("3⭐ Illaoi<!>")

	o.SortOutput()
	illaoi := o.Champ(1530, 50, 20).
		Shield(110, 400, 4).Heal(0, 60*4).
		Swing(4).Merge().
		Warden()
	illaoi.Simulate("2⭐ Illaoi")

	illaoi0 := o.Champ(2754, 50, 20).
		Shield(110, 400, 4).Heal(0, 70*4).
		Swing(4).Merge().
		Warden()
	illaoi0.Simulate("3⭐ Illaoi")

	amumu := o.Champ(1620, 50, 40).
		Heal(80, 300).Swing(3).
		Porcelain().
		Warden()
	amumu.Simulate("2⭐ Amumu")

	amumu0 := o.Champ(2916, 50, 40).
		Heal(80, 375).Swing(3).
		Porcelain().
		Warden()
	amumu0.Simulate("3⭐ Amumu")

	annie := o.Champ(1890, 50, 70).
		MixGrow(120, 500, 20).
		Swing(100).
		Annie()
	annie.Add(o.AS(12))
	annie.Simulate("2⭐ Annie")

	ornn := o.Champ(2160+150, 60+25, 120).
		MixShield(160, 350, 15, 4).Swing(2)
	ornn.Simulate("2⭐ Ornn")

	tahm := o.Champ(2754, 45, 40).
		MixShield(100, 0, 20, 3)
	tahm.Add(o.HpAmp(52))
	tahm.Simulate("3⭐ Tahm")

	tahm0 := o.Champ(1530, 45, 40).
		MixShield(100, 0, 20, 3)
	tahm0.Add(o.HpAmp(52))
	tahm0.Simulate("2⭐ Tahm")

	galio := o.Champ(2160, 60, 70).
		Shield(140, 200, 3).
		Buff(140, 4, o.AR(120+30)).Merge()
	galio.Add(o.HpAmp(40))
	galio.Simulate("2⭐ Galio")

	nasus := o.Champ(1800, 60, 70).
		Grow(150, 450).Buff(0, 100, o.AR(25)).Merge().
		//Shurima().
		Watcher()
	nasus.Simulate("2⭐ Nasus")
	o.OutputBySort()
}
