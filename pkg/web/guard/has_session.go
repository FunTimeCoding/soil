package guard

func (g *Mux) HasSession() bool {
	return g.session != nil
}
