package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"github.com/funtimecoding/soil/pkg/chromium/protocol/mock_protocol"
)

func (c *Client) Page(identifier string) face.ProtocolClient {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if p, found := c.pages[identifier]; found {
		return p
	}

	return mock_protocol.New()
}
