package parser

func (p *Parser) open(
	number int,
	kind string,
) {
	if p.blockStart == 0 {
		p.blockStart = number
		p.blockKind = kind
	}
}
