package server

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/face"
)

type Server struct {
	client   *mattermost.Client
	version  string
	reporter face.Reporter
}
