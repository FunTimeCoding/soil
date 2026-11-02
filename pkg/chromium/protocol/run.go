package protocol

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
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

		return zero, connection.New(
			errorsConstant.Timeout,
			constant.BrowserTab,
			"",
			fmt.Sprintf(constant.Asleep, p.timeout),
		)
	}
}
