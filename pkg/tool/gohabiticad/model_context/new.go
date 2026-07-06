package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	habitica "github.com/funtimecoding/soil/pkg/tool/gohabiticad/face"
)

func New(
	c habitica.HabiticaSource,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		habitica: c,
		reporter: r,
	}
	result.register()

	return result
}
