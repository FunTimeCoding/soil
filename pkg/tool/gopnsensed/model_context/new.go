package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	opnsense "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
)

func New(
	c opnsense.OpnsenseSource,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		opnsense: c,
		reporter: r,
	}
	result.register()

	return result
}
