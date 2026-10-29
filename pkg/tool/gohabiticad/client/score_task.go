package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ScoreTask(
	x context.Context,
	identifier string,
	direction string,
) (*client.ScoreResult, error) {
	r, e := c.client.ScoreTaskWithResponse(
		x,
		identifier,
		client.ScoreTaskParamsDirection(direction),
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("score task", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
