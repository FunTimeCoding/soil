package claim

import "github.com/funtimecoding/soil/pkg/tool/gomonitord/client"

func New(
	c *client.Client,
	updates chan Message,
) *Connection {
	return &Connection{client: c, updates: updates}
}
