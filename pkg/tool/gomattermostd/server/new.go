package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	mattermost "github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
)

func New(
	c mattermost.MattermostSource,
	version string,
	r face.Reporter,
) *Server {
	return &Server{client: c, version: version, reporter: r}
}
