package protocol

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	connectionConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"time"
)

func run[T any](
	p *Protocol,
	a chromedp.Action[T],
) (T, error) {
	if p.timeout <= 0 {
		return chromedp.Run(p.context, a)
	}

	var result T
	done := make(chan error, 1)
	go func() {
		v, e := chromedp.Run(p.context, a)
		result = v
		done <- e
	}()

	select {
	case e := <-done:
		return result, e
	case <-time.After(p.timeout):
		var zero T

		return zero, connection.New(
			connectionConstant.Timeout,
			constant.BrowserTab,
			"",
			fmt.Sprintf(constant.Asleep, p.timeout),
		)
	}
}
