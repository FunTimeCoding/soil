package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"time"
)

func (c *Client) Join(
	identifier string,
	turns int,
	seen time.Time,
) *target.Target {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result := target.New(identifier, identifier, seen, nil)
	result.TurnCount = turns
	c.sessions = append(c.sessions, result)

	return result
}
