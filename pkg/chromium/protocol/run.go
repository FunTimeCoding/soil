package protocol

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	connectionConstant "github.com/funtimecoding/soil/pkg/errors/constant"
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
		return connection.New(
			connectionConstant.Timeout,
			constant.BrowserTab,
			"",
			fmt.Sprintf(constant.Asleep, p.timeout),
		)
	}
}
