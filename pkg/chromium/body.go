package chromium

import (
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
)

func (c *Client) Body(identifier string) string {
	// Not canceling to leave the tab open
	o := c.TargetContext(identifier)
	var result string

	if false {
		c.RunContext(
			o,
			target.ActivateTarget(target.ID(identifier)), // Steals focus
			chromedp.OuterHTML(constant.BodySelector, &result),
		)
	}

	c.RunContext(o, chromedp.OuterHTML(constant.BodySelector, &result))

	return result
}
