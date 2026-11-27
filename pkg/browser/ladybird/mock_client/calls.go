package mock_client

import "github.com/funtimecoding/soil/pkg/browser/ladybird/types/call"

func (c *Client) Calls() []*call.Call {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return append([]*call.Call{}, c.calls...)
}
