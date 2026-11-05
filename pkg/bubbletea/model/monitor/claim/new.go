package claim

import "github.com/funtimecoding/soil/pkg/tool/gomonitord/face"

func New(
	c face.Client,
	updates chan Message,
) *Connection {
	return &Connection{client: c, updates: updates}
}
