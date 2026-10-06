package salt

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response/local_return"
)

func (c *Client) Highstate(
	target string,
) (map[string]local_return.LocalReturn, error) {
	return c.Local(target, constant.SaltHighstate, nil)
}
