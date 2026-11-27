package mock_client

import "github.com/chromedp/chromedp"

func (c *Client) Run(_ ...chromedp.Action[chromedp.Void]) error {
	return c.Failure
}
