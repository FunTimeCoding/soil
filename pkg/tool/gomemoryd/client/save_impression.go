package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"time"
)

func (c *RestClient) SaveImpression(
	x context.Context,
	content string,
	source string,
	capturedAt time.Time,
) error {
	at := capturedAt.UTC().Format(time.RFC3339)
	r, e := c.http.PostImpressionsWithResponse(
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
		return refusal.New("save impression", r.HTTPResponse, r.Body)
	}

	return nil
}
