package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"time"
)

func (c *Client) PostImpressions(
	x context.Context,
	content string,
	source string,
	capturedAt time.Time,
) error {
	at := capturedAt.UTC().Format(time.RFC3339)
	r, e := c.client.PostImpressionsWithResponse(
		x,
		client.PostImpressionsJSONRequestBody{
			Content:   content,
			Source:    &source,
			CreatedAt: &at,
		},
	)

	if e != nil {
		return e
	}

	if r.JSON200 == nil {
		return refusal.New("post impressions", r.HTTPResponse, r.Body)
	}

	return nil
}
