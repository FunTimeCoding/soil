package guard

import "net/http"

func (g *Mux) Open(
	pattern string,
	serve http.HandlerFunc,
) {
	g.mux.HandleFunc(pattern, serve)
}
