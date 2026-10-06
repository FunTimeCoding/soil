package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func postSave(
	c *Client,
	subject string,
	path string,
	body any,
) (*response.Save, error) {
	var out response.Save

	if e := c.basic.Post(path, body, &out); e != nil {
		return nil, e
	}

	if out.Result != constant.SavedResult {
		detail := formatValidation(out.Validation)

		if detail == "" {
			detail = out.Result
		}

		return nil, validation.New("%s rejected: %s", subject, detail)
	}

	return &out, nil
}
