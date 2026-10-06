package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	mattermost "github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
)

func New(
	c mattermost.MattermostSource,
	r face.Reporter,
) *Server {
	return &Server{client: c, reporter: r}
}
