package mock_parser

func New(
	tiddlers int,
	failure error,
) *Parser {
	return &Parser{tiddlers: tiddlers, failure: failure}
}
