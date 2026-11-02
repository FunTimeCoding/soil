package refusal_transport

import "net/http"

type Transport struct {
	host string
	next http.RoundTripper
}
