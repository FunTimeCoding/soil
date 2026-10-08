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

	signed := g.session(serve)
	g.mux.HandleFunc(
		pattern,
		func(
			w http.ResponseWriter,
			q *http.Request,
		) {
			if bearerAuthorized(q, g.tokens) {
				serve(w, q)

				return
			}

			signed(w, q)
		},
	)
}
