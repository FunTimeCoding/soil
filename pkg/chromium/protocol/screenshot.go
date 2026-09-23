package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Screenshot() ([]byte, error) {
	var result []byte
	e := chromedp.Run(p.context, chromedp.CaptureScreenshot(&result))

	if e != nil {
		return nil, e
	}

	return result, nil
}
