package tactics

func (g *Ground) Bruiser2() *Ground {
	g.Add(HpAmp(20))
	return g
}

func (g *Ground) Bruiser4() *Ground {
	g.Add(HpAmp(40))
	return g
}
