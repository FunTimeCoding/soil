package protocol

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
)

func (p *Protocol) Navigate(l string) error {
	return chromedp.Run(
		p.context,
		chromedp.Navigate(l),
		chromedp.WaitReady(constant.BodySelector),
	)
}
