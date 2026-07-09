package web

import "net/http"

func ServerPort(
	h http.Handler,
	port int,
) *http.Server {
	return Server(h, AddressPort(port))
}
