package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/provision/salt/constant"
)

func (c *Client) DeleteKey(minion string) ([]string, error) {
	b, e := c.Post(
		"",
		commandRequest{
			Client:   constant.WheelClient,
			Function: constant.KeyDelete,
			Match:    minion,
		},
	)

	if e != nil {
		return nil, e
	}

	var r response.Wheel

	if f := json.Unmarshal(b, &r); f != nil {
		return nil, f
	}

	if len(r.Return) == 0 || !r.Return[0].Result.Success {
		return nil, fmt.Errorf("failed to delete key for %s", minion)
	}

	return r.Return[0].Result.Affected.Minions, nil
}
