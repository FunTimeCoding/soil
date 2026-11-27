package protocol

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/browser/constant"
)

func (p *Protocol) Navigate(l string) error {
	return p.do(
		chromedp.Navigate(l),
		chromedp.WaitReady(constant.ChromiumBodySelector),
	)
}
