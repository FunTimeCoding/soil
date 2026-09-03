package guard

import "net/http"

func (g *Mux) OpenMount(
	pattern string,
	serve http.Handler,
) {
	g.Open(pattern, serve.ServeHTTP)
}
