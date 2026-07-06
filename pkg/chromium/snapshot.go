package chromium

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
)

func (c *Client) Snapshot(x context.Context) ([]*snapshot.Node, error) {
	var result []*snapshot.Node
	var fail error
	e := chromedp.Run(
		x,
		chromedp.ActionFunc(
			func(v context.Context) error {
				result, fail = snapshot.Take(v)

				return fail
			},
		),
	)

	if e != nil {
		return nil, e
	}

	return result, nil
}
