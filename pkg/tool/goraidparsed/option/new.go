package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/parser"
)

func New() *Parser {
	return &Parser{
		TopStats: func(path string) face.TopStatsParser {
			return parser.New(path)
		},
	}
}
