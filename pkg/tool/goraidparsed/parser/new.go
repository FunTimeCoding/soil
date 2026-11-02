package parser

func New(path string) *Parser {
	return &Parser{path: path}
}
