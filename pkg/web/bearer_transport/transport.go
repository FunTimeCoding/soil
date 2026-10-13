package bearer_transport

import "net/http"

type Transport struct {
	token string
	r     http.RoundTripper
}
