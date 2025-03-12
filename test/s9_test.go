package test

import (
	o "tactics"
	"testing"
)

func Test9(t *testing.T) {
	o.Level(0)
	o.SortOutput()

	nasus := o.Champ(1800, 60, 70).
		Grow(150, 450).Buff(0, 100, o.AR(25)).Merge().
		//Shurima().
		Watcher()
	nasus.Simulate("2⭐ Nasus")

	shen := o.Champ(1800, 60+25, 70).
		Shield(140, 450, 4).
		Invoker()
	shen.Add(o.HealAmp(40))
	shen.Simulate("2⭐ Shen")

	sejuani := o.Champ(1800, 60, 60).
		Shield(120, 700, 4)
	sejuani.Add(o.HpAmp(40))
	sejuani.Simulate("2⭐ Sejuani")
	o.OutputBySort()
}
