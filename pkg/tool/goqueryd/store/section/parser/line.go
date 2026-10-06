package parser

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section"
	"strings"
)

func (p *Parser) line(
	number int,
	text string,
) {
	switch {
	case p.inFront:
		p.open(number, constant.BlockFrontMatter)

		if number > 1 && text == constant.FrontMatterFence {
			p.flush(number)
			p.inFront = false
		}
	case p.inFence:
		if isFence(text) {
			p.flush(number)
			p.inFence = false
		}
	case isFence(text):
		p.flush(number - 1)
		p.open(number, constant.BlockCode)
		p.inFence = true
	case constant.AnyHeadingPattern.MatchString(text):
		p.closeSection(number - 1)
		p.current = section.NewSection(number)
		p.current.Level = len(text) - len(strings.TrimLeft(text, constant.HeadingMark))
		p.current.Title = text
		p.open(number, constant.BlockHeading)
		p.flush(number)
	case text == "":
		p.flush(number - 1)
	default:
		p.open(number, kind(text))
	}
}
