package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
)

func New(
	c *netbox.Client,
	s *store.Store,
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
		client:   c,
		store:    s,
		reporter: r,
	}
	result.register()

	return result
}
