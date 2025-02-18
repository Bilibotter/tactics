package tactics

import "fmt"

type shield_ struct {
	lastSec int
	health  int
}

func (g *Ground) addShield(shield ...*shield_) *Ground {
	g.shields = append(g.shields, shield...)
	if outputLevel >= 3 {
		total := 0
		for _, s := range shield {
			total += s.health
		}
		fmt.Printf("新增护盾值%d\n", total)
	}
	return g
}

func (s *shield_) Taken(dmg int) int {
	if !s.IsValid() {
		return dmg
	}
	if s.health >= dmg {
		s.health -= dmg
		return 0
	}
	s.health = 0
	return dmg - s.health
}

func (s *shield_) IsValid() bool {
	return s.lastSec > 0 && s.health > 0
}

func (s *shield_) handle(event Event, ground *Ground) {
	if event.Is(timeGoA) {
		s.lastSec--
	}
}
