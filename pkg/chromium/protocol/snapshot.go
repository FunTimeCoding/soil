package protocol

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
)

func (p *Protocol) Snapshot() ([]*snapshot.Node, error) {
	return run(
		p,
		func(
			v context.Context,
			t *chromedp.Target,
		) ([]*snapshot.Node, error) {
			return snapshot.Take(v, t)
		},
	)
}
