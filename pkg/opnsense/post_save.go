package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func postSave(
	c *Client,
	subject string,
	path string,
	body any,
	bare error,
) (*response.Save, error) {
	var out response.Save

	if e := c.basic.Post(path, body, &out); e != nil {
		return nil, e
	}

	if out.Result == constant.SavedResult {
		return &out, nil
	}

	if detail := formatValidation(out.Validation); detail != "" {
		return nil, validation.New("%s rejected: %s", subject, detail)
	}

	if bare != nil {
		return nil, bare
	}

	return nil, unexpected.Format(
		"%s save answered %s without validations",
		subject,
		out.Result,
	)
}
