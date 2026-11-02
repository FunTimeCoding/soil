package refusal_transport

import "net/http"

func New(
	host string,
	next http.RoundTripper,
) *Transport {
	return &Transport{host: host, next: next}
}
