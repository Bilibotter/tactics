package tactics

func (g *Ground) Heal(mana, heal int) *Ground {
	return g.MixHeal(mana, heal, 0, 0)
}

func (g *Ground) HealMax(mana, percent int) *Ground {
	return g.MixHeal(mana, 0, percent, 0)
}

func (g *Ground) Shield(mana, health, remain int) *Ground {
	return g.MixShield(mana, health, 0, remain)
}

func (g *Ground) MixHeal(mana, base, max, lose int) *Ground {
	skill := &skill_{
		call: make([]func(ground *Ground), 0, 1),
	}
	if base != 0 {
		skill.call = append(skill.call, healHp(base, true))
	}
	if max != 0 {
		skill.call = append(skill.call, healMax(max))
	}
	if lose != 0 {
		skill.call = append(skill.call, healLose(lose))
	}
	skill.mana = mana
	return g.addSkill(skill)
}

func (g *Ground) MixShield(mana, base, max, remain int) *Ground {
	skill := &skill_{
		call: make([]func(ground *Ground), 0, 1),
	}
	if base != 0 {
		skill.call = append(skill.call, shield(base, remain, true))
	}
	if max != 0 {
		skill.call = append(skill.call, shieldMax(max, remain, true))
	}
	skill.mana = mana
	return g.addSkill(skill)
}

func (g *Ground) Grow(mana, hp int) *Ground {
	skill := &skill_{
		call: []func(ground *Ground){growHp(hp, true)},
	}
	skill.mana = mana
	return g.addSkill(skill)
}

func (g *Ground) MixGrow(mana, base, max int) *Ground {
	skill := &skill_{
		call: []func(ground *Ground){},
	}
	// 先增长最大再增长固定
	if max != 0 {
		skill.call = append(skill.call, growMax(max))
	}
	if base != 0 {
		skill.call = append(skill.call, growHp(base, true))
	}
	skill.mana = mana
	return g.addSkill(skill)
}

func (g *Ground) Buff(mana, remain int, attrs ...*attrs_) *Ground {
	bf := buff(remain, attrs...)
	skill := &skill_{
		call: []func(ground *Ground){addLockBuff(bf)},
	}
	skill.mana = mana
	g.addSkill(skill)
	return g
}

func (g *Ground) Swing(duration int) *Ground {
	return g.Buff(0, duration).Merge()
}

func (g *Ground) Once() *Ground {
	return g.Buff(0, 1000).Merge()
}

func (g *Ground) Merge() *Ground {
	merged := &skill_{}
	merged.ground = g
	for head := g.skill; head != nil; head = head.next {
		merged.call = append(merged.call, head.call...)
		merged.mana += head.mana
		if merged.lock != nil {
			merged.lock = head.lock
		}
	}
	g.skill = merged
	return g
}

func (g *Ground) addSkill(skill *skill_) *Ground {
	skill.ground = g
	if g.skill == nil {
		g.skill = skill
		return g
	}
	tail := g.skill
	for ; tail.next != nil; tail = tail.next {
	}
	tail.next = skill
	return g
}
