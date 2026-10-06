package basic

import (
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) PostDiscard(path string) error {
	_, e := c.requester.Bytes(request.New(http.MethodPost, path))

	return e
}
