package guard

import "net/http"

func (g *Mux) SessionMount(
	pattern string,
	serve http.Handler,
) {
	g.Session(pattern, serve.ServeHTTP)
}
