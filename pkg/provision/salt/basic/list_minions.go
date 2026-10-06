package basic

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) ListMinions() ([]response.Minion, error) {
	var r response.MinionList

	if e := c.Get(constant.SaltMinionsPath, &r); e != nil {
		return nil, e
	}

	if len(r.Return) == 0 {
		return nil, nil
	}

	var result []response.Minion

	for _, v := range r.Return[0] {
		var m response.Minion

		if f := json.Unmarshal(v, &m); f != nil {
			continue
		}

		result = append(result, m)
	}

	return result, nil
}
