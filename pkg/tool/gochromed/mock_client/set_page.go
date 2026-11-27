package mock_client

import "github.com/funtimecoding/soil/pkg/browser/face"

func (c *Client) SetPage(
	identifier string,
	p face.ChromiumProtocolClient,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.pages[identifier] = p
}
