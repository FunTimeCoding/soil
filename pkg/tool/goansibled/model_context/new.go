package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/service"
)

func New(
	v *service.Service,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server:   server.New(constant.Identity, r).WithRecorder(t).Server(),
		service:  v,
		reporter: r,
	}
	result.register()

	return result
}
