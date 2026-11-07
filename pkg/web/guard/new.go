package guard

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func New(
	m *http.ServeMux,
	tokens []string,
) *Mux {
	result := &Mux{mux: m, tokens: tokens}

	if len(tokens) > 0 {
		result.Token(route.Get(constant.ReachablePath), reachable)
	}

	return result
}
