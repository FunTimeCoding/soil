package mock_client

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/protocol/mock_protocol"
	"github.com/funtimecoding/soil/pkg/browser/face"
)

func (c *Client) Page(identifier string) face.ChromiumProtocolClient {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if p, found := c.pages[identifier]; found {
		return p
	}

	return mock_protocol.New()
}
