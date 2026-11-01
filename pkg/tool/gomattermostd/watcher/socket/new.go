package socket

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(
	c face.Client,
	dispatch func(*model.WebSocketEvent),
) *Socket {
	return &Socket{client: c, dispatch: dispatch}
}
