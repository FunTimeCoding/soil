package guard

import "log"

func (g *Mux) requireTokens(pattern string) {
	if len(g.tokens) == 0 {
		log.Panicf("guard: protected route %s without tokens", pattern)
	}
}
