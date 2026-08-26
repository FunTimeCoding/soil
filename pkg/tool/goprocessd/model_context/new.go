package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	supervisor "github.com/funtimecoding/soil/pkg/tool/goprocessd/server"
)

func New(
	s *supervisor.Server,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		supervisor: s,
		reporter:   r,
	}
	result.register()

	return result
}
