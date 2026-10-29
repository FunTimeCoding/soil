package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	sublime "github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
)

func New(
	v sublime.Sublime,
	r face.Reporter,
) *Server {
	return &Server{service: v, reporter: r}
}
