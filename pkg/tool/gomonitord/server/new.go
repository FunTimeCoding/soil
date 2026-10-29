package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	monitor "github.com/funtimecoding/soil/pkg/tool/gomonitord/face"
)

func New(
	v monitor.Claims,
	r face.Reporter,
) *Server {
	return &Server{service: v, reporter: r}
}
