package mock_client

import "github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"

func (c *Client) NavigateOn(
	declaration string,
	events ...*message.Reply,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.navigations[declaration] = events
}
