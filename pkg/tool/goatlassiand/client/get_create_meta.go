package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetCreateMeta(
	x context.Context,
	project string,
	issueType string,
	expand string,
) ([]*client.CreateMetaField, error) {
	params := &client.GetCreateMetaParams{
		Project:   project,
		IssueType: issueType,
	}

	if expand != "" {
		params.Expand = &expand
	}

	r, e := c.client.GetCreateMetaWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("create meta", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
