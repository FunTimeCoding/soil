package web

import (
	"context"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
)

func Serve(
	c context.Context,
	r http.Handler,
	port int,
	verbose bool,
) {
	s := ServerPort(r, port)
	ServeAsynchronous(s)
	system.KillSignalBlock()
	GracefulShutdown(c, s, verbose)
}
