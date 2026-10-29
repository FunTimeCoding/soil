package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
	chromium "github.com/funtimecoding/soil/pkg/tool/gochromed/face"
)

func New(
	v chromium.Browser,
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
