package guard

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/route"
	"log"
	"net/http"
)

func New(
	m *http.ServeMux,
	tokens []string,
) *Mux {
	if len(tokens) == 0 {
		log.Panicf("guard: no tokens")
	}

	result := &Mux{mux: m, tokens: tokens}
	result.Token(route.Get(constant.ReachablePath), reachable)

	return result
}
