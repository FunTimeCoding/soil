package socket

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(
	c face.MattermostClient,
	dispatch func(*model.WebSocketEvent),
) *Socket {
	return &Socket{client: c, dispatch: dispatch}
}
