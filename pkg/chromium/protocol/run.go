package protocol

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

func (p *Protocol) run(actions ...chromedp.Action) error {
	if p.timeout <= 0 {
		return chromedp.Run(p.context, actions...)
	}

	done := make(chan error, 1)
	go func() {
		done <- chromedp.Run(p.context, actions...)
	}()

	select {
	case e := <-done:
		return e
	case <-time.After(p.timeout):
		return context.DeadlineExceeded
	}
}
