package basic

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/url"
)

func (c *Client) Post(
	path string,
	form url.Values,
) (json.RawMessage, error) {
	q := request.PostForm("", form)
	q.Locator = join.Empty(c.base, path)

	return c.send(q)
}
