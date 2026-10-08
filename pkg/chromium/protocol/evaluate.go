package protocol

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium"
)

func (p *Protocol) Evaluate(
	expression string,
	result any,
) error {
	value, e := run(p, chromedp.Evaluate[[]byte](expression))

	if e != nil {
		return e
	}

	return chromium.Decode(value, result)
}
