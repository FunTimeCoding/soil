package basic

import (
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) Delete(path string) error {
	_, e := c.requester.Bytes(request.New(http.MethodDelete, path))

	return e
}
