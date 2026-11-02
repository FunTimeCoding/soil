package chromium

import (
	"context"
	"github.com/chromedp/chromedp"
)

func (c *Client) RunContext(
	o context.Context,
	steps ...chromedp.Action[chromedp.Void],
) error {
	return chromedp.Do(o, steps...)
}
