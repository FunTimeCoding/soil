package salt

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) Highstate(
	target string,
) (map[string]response.LocalReturn, error) {
	return c.Local(target, constant.SaltHighstate, nil)
}
