package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Screenshot() ([]byte, error) {
	return run(p, chromedp.CaptureScreenshot())
}
