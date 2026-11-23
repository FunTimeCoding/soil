package graph

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (c *Client) Authorize(r *http.Request) error {
	t, e := c.current(false)

	if e != nil {
		return e
	}

	web.Bearer(r, t.AccessToken)

	return nil
}
