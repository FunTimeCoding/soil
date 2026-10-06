package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	iterm "github.com/funtimecoding/soil/pkg/tool/goitermd/face"
)

func New(
	c iterm.ItermSource,
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
