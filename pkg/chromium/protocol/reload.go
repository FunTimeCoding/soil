package protocol

import (
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
)

func (p *Protocol) Reload() {
	p.client.RunContext(p.context, page.Reload())
	p.client.RunContext(p.context, chromedp.WaitReady(constant.BodySelector))
}
