package protocol

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
)

func (p *Protocol) Snapshot() ([]*snapshot.Node, error) {
	var result []*snapshot.Node
	var fail error
	e := chromedp.Run(
		p.context,
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
