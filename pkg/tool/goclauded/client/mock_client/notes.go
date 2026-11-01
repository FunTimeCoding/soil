package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/notify_call"
	"slices"
)

func (c *Client) Notes() []*notify_call.Call {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return slices.Clone(c.notes)
}
