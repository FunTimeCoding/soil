package habitica

import (
	"github.com/funtimecoding/soil/pkg/habitica/user"
	"net/http"
)

func (c *Client) user() (*user.User, error) {
	var result *user.User
	r, e := c.do(http.MethodGet, "/user", nil)

	if e != nil {
		return result, e
	}

	return result, c.decode(r, &result)
}
