package server

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/face"
)

func New(
	c *mattermost.Client,
	version string,
	r face.Reporter,
) *Server {
	return &Server{
		client:   c,
		version:  version,
		reporter: r,
	}
}
