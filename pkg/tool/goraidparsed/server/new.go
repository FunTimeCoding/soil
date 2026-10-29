package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	parser "github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
)

func New(
	v parser.Generator,
	r face.Reporter,
) *Server {
	return &Server{service: v, reporter: r}
}
