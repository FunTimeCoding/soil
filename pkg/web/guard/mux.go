package guard

import "net/http"

type Mux struct {
	mux     *http.ServeMux
	tokens  []string
	session func(http.HandlerFunc) http.HandlerFunc
}
