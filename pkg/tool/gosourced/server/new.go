package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/diff"
)

func New(
	d *diff.Service,
	r face.Reporter,
) *Server {
	return &Server{diff: d, reporter: r}
}
