package basic

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/provision/types/salt_command_request"
)

func (c *Client) LocalClientAsync(
	target string,
	function string,
	arguments []string,
) (string, error) {
	var r response.Async

	if e := c.Post(
		"",
		salt_command_request.Request{
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
