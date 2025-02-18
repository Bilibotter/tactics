package tactics

func Champ(health int, armor int, mana int, attrs ...*attrs_) *Ground {
	champ := &champion_{
		Name:  "-",
		Shred: true,
		Wound: true,
		attrs_: attrs_{
			health:     (health + 600) * 128 / 100,
			armor:      armor + 60,
			speed:      100,
			healthAmp:  100,
			manaAmp:    100,
			recoverAmp: 100,
			dmgTaken:   100,
			ability:    100,
			castAmp:    100,
		},
		basicSpeed:  0.6,
		currentMana: mana,
	}
	for _, attr := range attrs {
		champ.Add(attr)
	}
	return &Ground{champion_: champ}
}

func (g *champion_) Passive(passives ...*passive) {
	for _, p := range passives {
		g.attach = append(g.attach, p)
	}
}
