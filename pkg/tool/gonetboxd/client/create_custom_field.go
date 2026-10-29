package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateCustomField(
	x context.Context,
	name string,
	label string,
	choiceSet string,
) (client.CustomField, error) {
	body := client.CreateCustomFieldRequest{Name: name, ChoiceSet: choiceSet}

	if label != "" {
		body.Label = &label
	}

	r, e := c.client.CreateCustomFieldWithResponse(x, body)

	if e != nil {
		return client.CustomField{}, e
	}

	if r.JSON201 == nil {
		return client.CustomField{}, refusal.New(
			"create custom field",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
