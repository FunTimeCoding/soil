package guard

import (
	"log"
	"net/http"
)

func (g *Mux) Session(
	pattern string,
	serve http.HandlerFunc,
) {
	if g.session == nil {
		log.Panicf("guard: session route %s without WithSession", pattern)
	}

	g.mux.HandleFunc(pattern, g.session(serve))
}
