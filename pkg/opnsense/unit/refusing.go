package unit

import (
	"net/http"
	"net/http/httptest"
)

func refusing(status int) *httptest.Server {
	return httptest.NewTLSServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.WriteHeader(status)
			},
		),
	)
}
