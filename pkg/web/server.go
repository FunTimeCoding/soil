package web

import "net/http"

func Server(
	h http.Handler,
	address string,
) *http.Server {
	return &http.Server{Addr: address, Handler: h}
}
