package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	sentry "github.com/funtimecoding/soil/pkg/tool/gosentryd/face"
)

func New(
	c sentry.SentrySource,
	organization string,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		client:       c,
		organization: organization,
		reporter:     r,
	}
	result.register()

	return result
}
