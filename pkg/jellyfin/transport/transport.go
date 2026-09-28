package transport

import "net/http"

type Transport struct {
	base  http.RoundTripper
	token string
}
