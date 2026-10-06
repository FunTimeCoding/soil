package opnsense

import (
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
	var out response.Save

	if e := c.basic.Post(path, struct{}{}, &out); e != nil {
		return e
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
