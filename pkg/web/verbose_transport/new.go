package verbose_transport

import "net/http"

func New(base http.RoundTripper) *Transport {
	return &Transport{base: base}
}
