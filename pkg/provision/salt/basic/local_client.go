package basic

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response/local_return"
	"github.com/funtimecoding/soil/pkg/provision/types/salt_command_request"
)

func (c *Client) LocalClient(
	target string,
	function string,
	arguments []string,
) (map[string]local_return.LocalReturn, error) {
	var r response.Local

	if e := c.Post(
		"",
		salt_command_request.Request{
			Client:     constant.SaltLocalClient,
			Target:     target,
			Function:   function,
			Arguments:  arguments,
			TargetType: constant.SaltGlobTarget,
			FullReturn: true,
		},
		&r,
	); e != nil {
		return nil, e
	}

	if len(r.Return) == 0 {
		return nil, nil
	}

	return r.Return[0], nil
}
