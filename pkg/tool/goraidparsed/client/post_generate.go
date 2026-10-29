package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"time"
)

func (c *Client) PostGenerate(
	x context.Context,
	files []string,
	date *time.Time,
) (client.GenerateResponse, error) {
	r, e := c.client.PostGenerateWithResponse(
		x,
		client.PostGenerateJSONRequestBody{Files: files, Date: date},
	)

	if e != nil {
		return client.GenerateResponse{}, e
	}

	if r.JSON200 == nil {
		return client.GenerateResponse{}, refusal.New(
			"generate report",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
