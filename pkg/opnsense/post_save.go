package opnsense

import (
	"encoding/json"
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
	b, e := c.basic.Post(path, body)

	if e != nil {
		return nil, e
	}

	var out response.Save

	if f := json.Unmarshal(b, &out); f != nil {
		return nil, f
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
