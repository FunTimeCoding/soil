package chromium

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) RunContext(
	o context.Context,
	steps ...chromedp.Action[chromedp.Void],
) {
	errors.PanicOnError(chromedp.Do(o, steps...))
}
