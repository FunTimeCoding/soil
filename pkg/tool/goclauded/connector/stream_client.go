package connector

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (c *Client) streamClient() *http.Client {
	if c.untrusted {
		return web.InsecureClient()
	}

	return web.Client()
}
