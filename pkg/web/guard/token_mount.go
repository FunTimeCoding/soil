package guard

import "net/http"

func (g *Mux) TokenMount(
	pattern string,
	serve http.Handler,
) {
	g.Token(pattern, serve.ServeHTTP)
}
