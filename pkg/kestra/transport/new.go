package transport

import "net/http"

func New(
	username string,
	password string,
	r http.RoundTripper,
) *Transport {
	return &Transport{username: username, password: password, r: r}
}
