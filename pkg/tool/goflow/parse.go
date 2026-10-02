package goflow

import (
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
)

func parse(source []byte) ast.Node {
	return parser.New(parser.WithExtensions(extension.NewTableParser())).
		Parse(source)
}
