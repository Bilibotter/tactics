package tactics

type skill_ struct {
	ground *Ground
	call   []func(ground *Ground)
	next   *skill_
	lock   *manaLock
	mana   int
}

func (s *skill_) costMana() int {
	g := s.ground
	reduce, shrink := g.manaReduce, g.manaShrink
	for _, a := range g.attach {
		if a.IsValid() {
			reduce += a.attr().manaReduce
			shrink += a.attr().manaShrink
		}
	}
	return (s.mana - reduce) * (100 - shrink) / 100
}

func (s *skill_) cast() {
	s.ground.filter(NewE(beforeCastA, s.ground.CastTimes))
	for _, call := range s.call {
		call(s.ground)
	}
	if s.lock != nil {
		lock := *s.lock
		s.ground.locks = append(s.ground.locks, &lock)
	}
	if s.next != nil {
		s.ground.skill = s.next
	}
	s.ground.filter(NewE(afterCastA, s.ground.CastTimes))
}

func (s *skill_) Loop() {
	tail := s
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = s
}
