package parser

import "strings"

func (p *Parser) closeSection(end int) {
	p.flush(end)

	if len(p.current.Blocks) == 0 {
		return
	}

	p.current.LastLine = end
	p.current.Text = strings.Join(p.lines[p.current.FirstLine-1:end], "\n")
	p.result = append(p.result, p.current)
}
