package basic

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/provision/salt/constant"
)

func (c *Client) LocalClient(
	target string,
	function string,
	arguments []string,
) (map[string]response.LocalReturn, error) {
	b, e := c.Post(
		"",
		commandRequest{
			Client:     constant.LocalClient,
			Target:     target,
			Function:   function,
			Arguments:  arguments,
			TargetType: constant.GlobTarget,
			FullReturn: true,
		},
	)

	if e != nil {
		return nil, e
	}

	var r response.Local

	if f := json.Unmarshal(b, &r); f != nil {
		return nil, f
	}

	if len(r.Return) == 0 {
		return nil, nil
	}

	return r.Return[0], nil
}
