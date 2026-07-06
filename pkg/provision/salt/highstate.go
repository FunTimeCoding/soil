package salt

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/provision/salt/constant"
)

func (c *Client) Highstate(
	target string,
) (map[string]response.LocalReturn, error) {
	return c.Local(target, constant.Highstate, nil)
}
