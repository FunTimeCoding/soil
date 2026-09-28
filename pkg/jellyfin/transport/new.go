package transport

import "net/http"

func New(token string) *Transport {
	return &Transport{token: token, base: http.DefaultTransport}
}
