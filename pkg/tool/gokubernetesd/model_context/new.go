package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
)

func New(
	s *service.Service,
	readOnly bool,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithResources().WithRecorder(t).Server(),
		service:  s,
		readOnly: readOnly,
		reporter: r,
	}
	result.register()
	result.registerResources()

	if !readOnly {
		result.registerWrite()
	}

	return result
}
