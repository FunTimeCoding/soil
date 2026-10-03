package protocol

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
)

func (p *Protocol) Navigate(l string) error {
	return p.run(
		chromedp.Navigate(l),
		chromedp.WaitReady(constant.BodySelector),
	)
}
