package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/check"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/diff"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/usage"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
)

func New(
	c *usage.Service,
	k *check.Service,
	d *diff.Service,
	m *module.Service,
	w *workspace.Service,
	r face.Reporter,
) *Server {
	return &Server{
		usage:     c,
		check:     k,
		diff:      d,
		module:    m,
		workspace: w,
		reporter:  r,
	}
}
