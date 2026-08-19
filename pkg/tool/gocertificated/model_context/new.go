package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
)

func New(
	s *store.Store,
	v *service.Service,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server:   server.New(constant.Identity, version).WithRecorder(t).Server(),
		store:    s,
		service:  v,
		reporter: r,
	}
	result.register()

	return result
}
