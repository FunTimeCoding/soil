package mock_client

import "github.com/funtimecoding/soil/pkg/chromium/face"

func (c *Client) SetPage(
	identifier string,
	p face.ProtocolClient,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.pages[identifier] = p
}
