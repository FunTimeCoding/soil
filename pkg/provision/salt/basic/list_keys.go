package basic

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) ListKeys() (*response.KeysReturn, error) {
	b, e := c.Get(constant.SaltKeysPath)

	if e != nil {
		return nil, e
	}

	var r response.Keys

	if f := json.Unmarshal(b, &r); f != nil {
		return nil, f
	}

	return &r.Return, nil
}
