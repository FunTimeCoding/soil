package ladybird

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (c *Client) driverSend(q *http.Request) *http.Response {
	q.Close = true

	return web.Send(c.pool.Client(), q)
}
