package web

import "github.com/funtimecoding/soil/pkg/web/locator"

func Link(
	host string,
	port int,
	secure bool,
) string {
	return locator.NewEndpoint(host, port, secure).String()
}
