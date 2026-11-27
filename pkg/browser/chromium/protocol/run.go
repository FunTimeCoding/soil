package protocol

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"time"
)

func run[T any](
	p *Protocol,
	a chromedp.Action[T],
) (T, error) {
	x, f := p.target()

	if f != nil {
		var zero T

		return zero, f
	}

	if p.timeout <= 0 {
		return chromedp.Run(x, a)
	}

	var result T
	done := make(chan error, 1)
	go func() {
		v, e := chromedp.Run(x, a)
		result = v
		done <- e
	}()

	select {
	case e := <-done:
		return result, e
	case <-time.After(p.timeout):
		var zero T

		return zero, timeout.Format(
			"%s: %s",
			constant.ChromiumBrowserTab,
			fmt.Sprintf(constant.ChromiumAsleep, p.timeout),
		)
	}
}
