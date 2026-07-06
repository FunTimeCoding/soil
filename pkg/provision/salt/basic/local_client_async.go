package basic

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/provision/salt/constant"
)

func (c *Client) LocalClientAsync(
	target string,
	function string,
	arguments []string,
) (string, error) {
	b, e := c.Post(
		"",
		commandRequest{
			Client:     constant.LocalAsyncClient,
			Target:     target,
			Function:   function,
			Arguments:  arguments,
			TargetType: constant.GlobTarget,
		},
	)

	if e != nil {
		return "", e
	}

	var r response.Async

	if f := json.Unmarshal(b, &r); f != nil {
		return "", f
	}

	if len(r.Return) == 0 {
		return "", nil
	}

	return r.Return[0].JID, nil
}
