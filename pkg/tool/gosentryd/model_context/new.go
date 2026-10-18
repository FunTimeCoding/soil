package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/service"
)

func New(
	v *service.Service,
	host string,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server:   server.New(constant.Identity).WithRecorder(t).Server(),
		service:  v,
		host:     host,
		reporter: r,
		reacher:  reacher.New(),
	}
	result.register()

	return result
}
