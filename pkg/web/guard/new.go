package guard

import (
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

	return &Mux{mux: m, tokens: tokens}
}
