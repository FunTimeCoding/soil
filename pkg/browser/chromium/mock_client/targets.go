package mock_client

import "github.com/chromedp/cdproto/target"

func (c *Client) Targets() ([]*target.Info, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return []*target.Info{}, nil
}
