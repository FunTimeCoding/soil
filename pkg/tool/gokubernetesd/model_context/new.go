package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
)

func New(
	s *service.Service,
	armed bool,
	r face.Reporter,
	t face.Recorder,
) *Server {
	chosen := selection.New()
	result := &Server{
		server: server.New(constant.Identity, r).WithResources().WithRecorder(
			t,
		).WithSessionEnd(chosen.Clear).Server(),
		service:   s,
		selection: chosen,
		armed:     armed,
		reporter:  r,
	}
	result.register()
	result.registerResources()

	if armed {
		result.registerWrite()
	}

	return result
}
