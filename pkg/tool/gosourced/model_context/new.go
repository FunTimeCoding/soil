package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
)

func New(
	s *service.Service,
	r face.Reporter,
	t face.Recorder,
) *Server {
	chosen := selection.New()
	result := &Server{
		server: server.New(constant.Identity).WithRecorder(t).WithSessionEnd(
			chosen.Clear,
		).Server(),
		service:   s,
		selection: chosen,
		reporter:  r,
	}
	result.register()

	return result
}
