package bearer_transport

import "net/http"

func New(
	token string,
	r http.RoundTripper,
) *Transport {
	return &Transport{token: token, r: r}
}
