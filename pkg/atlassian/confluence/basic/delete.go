package basic

import (
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) Delete(l string) error {
	q := request.Absolute(l)
	q.Method = http.MethodDelete
	_, e := c.requester.Bytes(q)

	return e
}
