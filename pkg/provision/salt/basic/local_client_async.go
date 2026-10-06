package basic

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) LocalClientAsync(
	target string,
	function string,
	arguments []string,
) (string, error) {
	var r response.Async

	if e := c.Post(
		"",
		commandRequest{
			Client:     constant.SaltLocalAsyncClient,
			Target:     target,
			Function:   function,
			Arguments:  arguments,
			TargetType: constant.SaltGlobTarget,
		},
		&r,
	); e != nil {
		return "", e
	}

	if len(r.Return) == 0 {
		return "", nil
	}

	return r.Return[0].JID, nil
}
