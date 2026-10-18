package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/impression_call"
	"time"
)

func (c *Client) SaveImpression(
	_ context.Context,
	content string,
	source string,
	capturedAt time.Time,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	c.Impressions = append(
		c.Impressions,
		impression_call.Call{
			Content:    content,
			Source:     source,
			CapturedAt: capturedAt,
		},
	)

	return nil
}
