package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListProjects(x context.Context) ([]*client.JiraProject, error) {
	r, e := c.client.ListProjectsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("projects", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
