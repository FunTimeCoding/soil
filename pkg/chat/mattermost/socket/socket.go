package socket

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/mattermost/mattermost/server/public/model"
)

type Socket struct {
	client   face.Client
	dispatch func(*model.WebSocketEvent)
	opened   bool
}
