package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Location() (string, error) {
	var result string

	if e := chromedp.Run(p.context, chromedp.Location(&result)); e != nil {
		return "", e
	}

	return result, nil
}
