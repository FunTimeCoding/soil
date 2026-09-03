package guard

import "net/http"

func (g *Mux) WithSession(
	session func(http.HandlerFunc) http.HandlerFunc,
) *Mux {
	g.session = session

	return g
}
