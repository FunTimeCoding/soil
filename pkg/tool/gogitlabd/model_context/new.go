package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	forge "github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
)

func New(
	c forge.Forge,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server:   server.New(constant.Identity).WithRecorder(t).Server(),
		client:   c,
		reporter: r,
	}
	result.register()

	return result
}
