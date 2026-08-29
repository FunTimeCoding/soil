package opnsense

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func postDelete(
	c *Client,
	subject string,
	path string,
	identifier string,
) error {
	b, e := c.basic.Post(path, struct{}{})

	if e != nil {
		return e
	}

	var out response.Save

	if f := json.Unmarshal(b, &out); f != nil {
		return f
	}

	if out.Result == constant.NotFoundResult {
		return not_found.New(subject, identifier)
	}

	if out.Result != constant.DeletedResult {
		return unexpected.Format(
			"unexpected %s delete: %s",
			subject,
			out.Result,
		)
	}

	return nil
}
