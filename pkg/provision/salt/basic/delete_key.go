package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) DeleteKey(minion string) ([]string, error) {
	var r response.Wheel

	if e := c.Post(
		"",
		commandRequest{
			Client:   constant.SaltWheelClient,
			Function: constant.SaltKeyDelete,
			Match:    minion,
		},
		&r,
	); e != nil {
		return nil, e
	}

	if len(r.Return) == 0 || !r.Return[0].Result.Success {
		return nil, unexpected.Format("delete key: %s", minion)
	}

	return r.Return[0].Result.Affected.Minions, nil
}
