package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/census"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/check"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/diff"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
)

func New(
	c *census.Service,
	k *check.Service,
	d *diff.Service,
	m *module.Service,
	w *workspace.Service,
	r face.Reporter,
) *Server {
	return &Server{
		census:    c,
		check:     k,
		diff:      d,
		module:    m,
		workspace: w,
		reporter:  r,
	}
}
