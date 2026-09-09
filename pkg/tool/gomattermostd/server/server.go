package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	mattermost "github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
)

type Server struct {
	client   mattermost.MattermostSource
	version  string
	reporter face.Reporter
}
