package protocol

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (p *Protocol) MustFindNode(
	s string,
	index int,
) (*chromedp.Node, bool) {
	result, okay, e := p.FindNode(s, index)
	errors.PanicOnError(e)

	return result, okay
}
