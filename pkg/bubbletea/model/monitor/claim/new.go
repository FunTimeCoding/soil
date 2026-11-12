package claim

import (
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim/message"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/face"
)

func New(
	c face.Client,
	updates chan message.Message,
) *Connection {
	return &Connection{client: c, updates: updates}
}
