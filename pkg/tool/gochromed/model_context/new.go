package model_context

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
)

func New(
	c *chromium.Client,
	downloadDirectory string,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		client:            c,
		downloadDirectory: downloadDirectory,
		snapshotCache:     make(map[string]map[string]int64),
		reporter:          r,
	}
	result.register()

	return result
}
