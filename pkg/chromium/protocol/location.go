package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Location() (string, error) {
	var result string

	if e := p.run(chromedp.Location(&result)); e != nil {
		return "", e
	}

	return result, nil
}
