package web_client

import (
	"github.com/funtimecoding/soil/pkg/web/web_client/web_response"
	"net/http"
	"time"
)

func (c *Client) Get(locator string) (*web_response.Response, error) {
	start := c.clock.Now()
	response, e := http.Get(locator)

	return web_response.New(response, time.Since(start).Milliseconds()), e
}
