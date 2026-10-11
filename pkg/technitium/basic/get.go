package basic

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Get(path string) (json.RawMessage, error) {
	return c.send(request.Absolute(join.Empty(c.base, path)))
}
