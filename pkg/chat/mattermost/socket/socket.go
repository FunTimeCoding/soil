package socket

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/mattermost/mattermost/server/public/model"
)

type Socket struct {
	client   face.MattermostClient
	dispatch func(*model.WebSocketEvent)
	opened   bool
}
