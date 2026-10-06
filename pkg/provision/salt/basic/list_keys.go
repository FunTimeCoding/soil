package basic

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) ListKeys() (*response.KeysReturn, error) {
	var r response.Keys

	if e := c.Get(constant.SaltKeysPath, &r); e != nil {
		return nil, e
	}

	return &r.Return, nil
}
