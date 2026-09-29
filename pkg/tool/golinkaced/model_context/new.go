package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	linkace "github.com/funtimecoding/soil/pkg/tool/golinkaced/face"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/service"
)

func New(
	c linkace.LinkAceSource,
	s *service.Service,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		client:   c,
		service:  s,
		reporter: r,
	}
	result.register()

	return result
}
