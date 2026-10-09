package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/reacher"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	sentry "github.com/funtimecoding/soil/pkg/tool/gosentryd/face"
)

func New(
	c sentry.SentrySource,
	organization string,
	host string,
	r face.Reporter,
	t face.Recorder,
) *Server {
	result := &Server{
		server:       server.New(constant.Identity).WithRecorder(t).Server(),
		client:       c,
		organization: organization,
		host:         host,
		reporter:     r,
		reacher:      reacher.New(),
	}
	result.register()

	return result
}
