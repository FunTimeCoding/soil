package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/constant"
	sublime "github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
)

func New(
	v sublime.Sublime,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server:   server.New(constant.Identity).WithRecorder(t).Server(),
		service:  v,
		reporter: r,
	}
	result.register()

	return result
}
