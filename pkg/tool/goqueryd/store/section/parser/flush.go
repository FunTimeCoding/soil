package parser

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section"
	"strings"
)

func (p *Parser) flush(end int) {
	if p.blockStart == 0 || end < p.blockStart {
		p.blockStart = 0

		return
	}

	p.current.Blocks = append(
		p.current.Blocks,
		section.NewBlock(
			p.blockKind,
			p.blockStart,
			end,
			strings.Join(p.lines[p.blockStart-1:end], "\n"),
		))
	p.blockStart = 0
}
