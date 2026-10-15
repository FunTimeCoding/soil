package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateCustomField(
	name string,
	label string,
	choiceSet string,
) *response.Response {
	body := client.CreateCustomFieldRequest{Name: name, ChoiceSet: choiceSet}

	if label != "" {
		body.Label = &label
	}

	result, e := c.client.CreateCustomField(c.context, body)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
