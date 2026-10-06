package basic

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response/local_return"
)

func (c *Client) LocalClient(
	target string,
	function string,
	arguments []string,
) (map[string]local_return.LocalReturn, error) {
	var r response.Local

	if e := c.Post(
		"",
		commandRequest{
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
