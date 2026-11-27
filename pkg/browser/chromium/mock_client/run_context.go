package mock_client

import (
	"context"
	"github.com/chromedp/chromedp"
)

func (c *Client) RunContext(
	_ context.Context,
	_ ...chromedp.Action[chromedp.Void],
) error {
	return c.Failure
}
