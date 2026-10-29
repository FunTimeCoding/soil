package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) LinkIssues(
	x context.Context,
	key string,
	target string,
	linkType string,
) error {
	body := client.LinkIssuesJSONRequestBody{TargetKey: target}

	if linkType != "" {
		body.LinkType = &linkType
	}

	r, e := c.client.LinkIssuesWithResponse(x, key, body)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("link issues", r.HTTPResponse, r.Body)
	}

	return nil
}
