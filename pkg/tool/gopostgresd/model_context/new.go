package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/mark/server"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/service"
)

func New(
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
		service:  s,
		reporter: r,
	}
	result.register()

	return result
}
