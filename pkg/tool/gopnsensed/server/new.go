package server

import (
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
)

func New(
	c face.OpnsenseSource,
	r library.Reporter,
) *Server {
	return &Server{opnsense: c, reporter: r}
}
