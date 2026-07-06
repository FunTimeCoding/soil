package web

import "github.com/funtimecoding/soil/pkg/web/locator"

func Link(
	host string,
	port int,
	secure bool,
) string {
	l := locator.New(host).Port(port)

	if !secure {
		l = l.Insecure()
	}

	return l.String()
}
