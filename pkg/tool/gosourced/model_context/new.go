package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/file"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/packages"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol"
)

func New(
	f *file.Service,
	m *module.Service,
	p *packages.Service,
	a *pattern.Service,
	e *reference.Service,
	y *symbol.Service,
	r face.Reporter,
	t face.Recorder,
) *Server {
	chosen := selection.New()
	result := &Server{
		server: server.New(constant.Identity, r).WithRecorder(t).WithSessionEnd(
			chosen.Clear,
		).Server(),
		file:      f,
		module:    m,
		packages:  p,
		pattern:   a,
		reference: e,
		symbol:    y,
		selection: chosen,
		reporter:  r,
	}
	result.register()

	return result
}
