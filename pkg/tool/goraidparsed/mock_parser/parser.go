package mock_parser

type Parser struct {
	tiddlers int
	failure  error
	calls    int
	workdir  string
}
