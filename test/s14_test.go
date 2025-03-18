package test

import (
	o "tactics"
	"testing"
)

// 同等承伤时长
// 治疗(不一定有减疗) > 加血(对抗百分比不行) > 护盾(护盾不一定能完全消耗)
// > 减伤(真伤) > 双抗(真伤)

// 双肉羁绊都是城墙

func Test14V0(t *testing.T) {
	o.SortOutput()
	o.Level(0)
	// 晕两次，一次2秒
	sejuani := o.Champ(1800, 60, 60)
	sejuani.Heal(150, 0)
	sejuani.Add(o.ArAmp(30))
	sejuani.Bastion(4)
	sejuani.Simulate("Sejuani")

	neeko := o.Champ(1800, 60, 30)
	neeko.MixShield(95, 300, 10, 4)
	neeko.Strategist(5)
	neeko.Simulate("Neeko[5]")

	neeko = o.Champ(1800, 60, 30)
	neeko.MixShield(95, 300, 10, 4)
	neeko.Strategist(4)
	neeko.Simulate("Neeko[4]")

	// 晕2次
	leona := o.Champ(1980, 60, 50)
	leona.Buff(110, 4, o.DR(60))
	leona.Animal(7)
	leona.Vanguard0(2)
	leona.Simulate("Leona[7]")

	// 晕3次
	leona = o.Champ(1980, 60, 50)
	leona.Buff(110, 4, o.DR(60))
	leona.Vanguard0(4)
	leona.Simulate("Leona[4]")

	chogath := o.Champ(1800, 50, 60)
	chogath.Wound = false // 将加血改成治疗
	chogath.MixHeal(120, 160, 10, 0)
	chogath.Bruiser(4)
	chogath.Simulate("chogath[4]")

	chogath = o.Champ(1800, 50, 60)
	chogath.Wound = false // 将加血改成治疗
	chogath.MixHeal(120, 160, 10, 0)
	chogath.Bruiser(6)
	chogath.Simulate("chogath[6]")

	mordekaiser := o.Champ(1530, 50, 30)
	mordekaiser.Shield(80, 525, 4)
	mordekaiser.Bruiser(4)
	mordekaiser.Simulate("Mordekaiser[2*]")

	mordekaiser = o.Champ(2754, 50, 30)
	mordekaiser.Shield(80, 700, 4) // swing0
	mordekaiser.Bruiser(4)
	mordekaiser.Simulate("Mordekaiser[3*]")

	jarvan := o.Champ(1530, 50, 30)
	jarvan.Shield(90, 300+50*2, 4) // 0.4swing
	jarvan.Vanguard0(4)
	jarvan.Simulate("Jarvan[2*]")

	jarvan = o.Champ(2754, 50, 30)
	jarvan.Shield(90, 350+80*2, 4)
	jarvan.Vanguard0(4)
	jarvan.Simulate("Jarvan[3*]")

	gragas := o.Champ(1530, 50, 30)
	gragas.MixHeal(90, 325, 10, 0).Swing(1) // swing 1.3
	gragas.Bruiser(4)
	gragas.Simulate("Gragas[2*]")

	gragas = o.Champ(2754, 50, 30)
	gragas.MixHeal(90, 350, 10, 0).Swing(1)
	gragas.Bruiser(4)
	gragas.Simulate("Gragas[3*]")

	galio := o.Champ(1620, 50, 20)
	galio.Buff(80, 3, o.DR(55)).Swing(1) // swing 0.4
	galio.Bastion(4)
	galio.Simulate("Galio[2*]")

	galio = o.Champ(2916, 50, 20)
	galio.Buff(80, 3, o.DR(60)).Swing(1)
	galio.Bastion(4)
	galio.Simulate("Galio[2*]")

	braum := o.Champ(1520, 50, 30)
	braum.MixShield(100, 475, 10, 4) // 0.5swing
	braum.Vanguard0(4)
	braum.Simulate("Braum[2*]")

	braum = o.Champ(2754, 50, 30)
	braum.MixShield(100, 500, 10, 4)
	braum.Vanguard0(4)
	braum.Simulate("Braum[3*]")

	skarnar := o.Champ(2592, 50, 25)
	skarnar.Shield(75, 450, 3) // swing0.2
	skarnar.Vanguard0(4)
	skarnar.Simulate("Skanar[3*]")

	rhaast := o.Champ(2268, 40, 50)
	rhaast.MixHeal(100, 170, 10, 0).Swing(1)
	rhaast.Vanguard0(4)
	rhaast.Simulate("Rhaast[3*]")

	illaoi := o.Champ(2592, 45, 0)
	illaoi.Heal(60, 500).Swing(2)
	illaoi.Bastion(4)
	illaoi.Simulate("Illaoi[3*]")

	ekko := o.Champ(2592, 45, 30)
	ekko.Heal(80, 425).Swing(1)
	ekko.Strategist(4)
	ekko.Simulate("Ekko[3*]")

	// 有大头目
	darius := o.Champ(2592, 45, 30)
	darius.Syndicate(5)
	darius.Bruiser(2)
	darius.MixHeal(80, 240+50*2.5, 5+2*2.5, 0).Swing(1)
	darius.Simulate("Darius[3*]")

	vi := o.Champ(2108, 40, 30)
	// 急速衰减
	vi.MixShield(70, 425*7/10, 15*7/10, 4)
	vi.Vanguard0(4)
	vi.Simulate("Vi[3*]")

	jax := o.Champ(2108, 40, 20)
	jax.Shield(80, 525, 4)
	jax.Bastion(4)
	jax.Simulate("Jax[3*]")

	alistar := o.Champ(2106, 40, 20)
	alistar.Add(o.BLK(25))
	alistar.Bruiser(4)
	alistar.Simulate("Alistar[3*]")

	poppy := o.Champ(2106, 40, 20)
	poppy.Shield(70, 525, 4)
	poppy.Bastion(4)
	poppy.Simulate("Poppy[3*]")

	o.OutputBySort()
}

func Test14(t *testing.T) {
	o.SortOutput()
	o.Level(0)
	// 晕两次，一次2秒
	sejuani := o.Champ(1800, 60, 60)
	sejuani.Heal(150, 0)
	sejuani.Add(o.ArAmp(30))
	sejuani.Bastion(4)
	sejuani.Simulate("Sejuani")

	neeko := o.Champ(1800, 60, 30)
	neeko.MixShield(95, 300, 10, 4)
	neeko.Strategist(5)
	neeko.Simulate("Neeko[5]")

	neeko = o.Champ(1800, 60, 30)
	neeko.MixShield(95, 300, 10, 4)
	neeko.Strategist(4)
	neeko.Simulate("Neeko[4]")

	// 晕2次
	leona := o.Champ(1980, 60, 50)
	leona.Buff(110, 4, o.DR(60))
	leona.Animal(7)
	leona.Vanguard(2)
	leona.Simulate("Leona[7]")

	// 晕3次
	leona = o.Champ(1980, 60, 50)
	leona.Buff(110, 4, o.DR(60))
	leona.Vanguard(4)
	leona.Simulate("Leona[4]")

	chogath := o.Champ(1800, 50, 60)
	chogath.Wound = false // 将加血改成治疗
	chogath.MixHeal(120, 160, 10, 0)
	chogath.Bruiser(4)
	chogath.Simulate("chogath[4]")

	chogath = o.Champ(1800, 50, 60)
	chogath.Wound = false // 将加血改成治疗
	chogath.MixHeal(120, 160, 10, 0)
	chogath.Bruiser(6)
	chogath.Simulate("chogath[6]")

	mordekaiser := o.Champ(1530, 50, 30)
	mordekaiser.Shield(80, 525, 4)
	mordekaiser.Bruiser(4)
	mordekaiser.Simulate("Mordekaiser[2*]")

	mordekaiser = o.Champ(2754, 50, 30)
	mordekaiser.Shield(80, 700, 4) // swing 0
	mordekaiser.Bruiser(4)
	mordekaiser.Simulate("Mordekaiser[3*]")

	jarvan := o.Champ(1530, 50, 30)
	jarvan.Shield(90, 300+50*2, 4) // 0.4swing
	jarvan.Vanguard(4)
	jarvan.Simulate("Jarvan[2*]")

	jarvan = o.Champ(2754, 50, 30)
	jarvan.Shield(90, 350+80*2, 4)
	jarvan.Vanguard(4)
	jarvan.Simulate("Jarvan[3*]")

	gragas := o.Champ(1530, 50, 30)
	gragas.MixHeal(90, 325, 10, 0).Swing(1) // swing 1.3
	gragas.Bruiser(4)
	gragas.Simulate("Gragas[2*]")

	gragas = o.Champ(2754, 50, 30)
	gragas.MixHeal(90, 350, 10, 0).Swing(1)
	gragas.Bruiser(4)
	gragas.Simulate("Gragas[3*]")

	galio := o.Champ(1620, 50, 20)
	galio.Buff(80, 3, o.DR(55)).Swing(1) // swing 0.4
	galio.Bastion(4)
	galio.Simulate("Galio[3*]")

	galio = o.Champ(2916, 50, 20)
	galio.Buff(80, 3, o.DR(60)).Swing(1)
	galio.Bastion(4)
	galio.Simulate("Galio[2*]")

	braum := o.Champ(1520, 50, 30)
	braum.MixShield(100, 475, 10, 4) // 0.5swing
	braum.Vanguard(4)
	braum.Simulate("Braum[2*]")

	braum = o.Champ(2754, 50, 30)
	braum.MixShield(100, 500, 10, 4)
	braum.Vanguard(4)
	braum.Simulate("Braum[3*]")

	skarnar := o.Champ(2592, 50, 25)
	skarnar.Shield(75, 450, 3) // swing 0.2
	skarnar.Vanguard(4)
	skarnar.Simulate("Skanar[3*]")

	rhaast := o.Champ(2268, 40, 50)
	rhaast.MixHeal(100, 170, 10, 0).Swing(1)
	rhaast.Vanguard0(4)
	rhaast.Simulate("Rhaast[3*]")

	illaoi := o.Champ(2592, 45, 0)
	illaoi.Heal(60, 500).Swing(2)
	illaoi.Bastion(4)
	illaoi.Simulate("Illaoi[3*]")

	ekko := o.Champ(2592, 45, 30)
	ekko.Heal(80, 425).Swing(1)
	ekko.Strategist(4)
	ekko.Simulate("Ekko[3*]")

	// 有大头目
	darius := o.Champ(2592, 45, 30)
	darius.Syndicate(5)
	darius.Bruiser(2)
	darius.MixHeal(80, 240+50*2.5, 5+2*2.5, 0).Swing(1)
	darius.Simulate("Darius[3*]")

	vi := o.Champ(2108, 40, 30)
	vi.MixShield(70, 425*7/10, 15*7/10, 4)
	vi.Vanguard(4)
	vi.Simulate("Vi[3*]")

	jax := o.Champ(2108, 40, 20)
	jax.Shield(80, 525, 4)
	jax.Bastion(4)
	jax.Simulate("Jax[3*]")

	alistar := o.Champ(2106, 40, 20)
	alistar.Add(o.BLK(25))
	alistar.Bruiser(4)
	alistar.Simulate("Alistar[3*]")

	poppy := o.Champ(2106, 40, 20)
	poppy.Shield(70, 525, 4)
	poppy.Bastion(4)
	poppy.Simulate("Poppy[3*]")

	o.OutputBySort()
}

func Test14Single(t *testing.T) {
	o.Level(3)
	skarnar := o.Champ(2592, 50, 25)
	skarnar.Shield(75, 450, 3) // swing0.2
	skarnar.Vanguard0(4)
	skarnar.Simulate("Skanar[3*]")
}
