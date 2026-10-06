package salt

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response/local_return"

func (c *Client) Local(
	glob string,
	function string,
	a []string,
) (map[string]local_return.LocalReturn, error) {
	return c.basic.LocalClient(glob, function, a)
}
