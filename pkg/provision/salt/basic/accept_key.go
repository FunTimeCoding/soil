package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/provision/types/salt_command_request"
)

func (c *Client) AcceptKey(minion string) ([]string, error) {
	var r response.Wheel

	if e := c.Post(
		"",
		salt_command_request.Request{
			Client:   constant.SaltWheelClient,
			Function: constant.SaltKeyAccept,
			Match:    minion,
		},
		&r,
	); e != nil {
		return nil, e
	}

	if len(r.Return) == 0 || !r.Return[0].Result.Success {
		return nil, unexpected.Format("accept key: %s", minion)
	}

	return r.Return[0].Result.Affected.Minions, nil
}
