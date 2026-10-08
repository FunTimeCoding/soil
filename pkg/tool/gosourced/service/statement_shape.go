package service

import (
	"go/ast"
	"go/token"
	"strings"
)

func statementShape(
	content []byte,
	set *token.FileSet,
	node ast.Node,
	anchor ast.Node,
) (string, string) {
	start := set.Position(node.Pos()).Offset
	end := set.Position(node.End()).Offset
	lineEnd := end

	if index := strings.IndexByte(
		string(content[start:end]),
		'\n',
	); index >= 0 {
		lineEnd = start + index
	}

	var b strings.Builder
	cursor := start
	ast.Inspect(
		node,
		func(n ast.Node) bool {
			if n == nil {
				return false
			}

			if n.Pos() >= anchor.Pos() && n.End() <= anchor.End() {
				return false
			}

			from := set.Position(n.Pos()).Offset
			to := set.Position(n.End()).Offset

			if to > lineEnd {
				return true
			}

			var text string

			switch leaf := n.(type) {
			case *ast.Ident:
				text = "IDENT"
			case *ast.BasicLit:
				text = leaf.Kind.String()
			default:
				return true
			}

			b.Write(content[cursor:from])
			b.WriteString(text)
			cursor = to

			return true
		},
	)
	b.Write(content[cursor:lineEnd])

	return strings.TrimSpace(b.String()),
		strings.TrimSpace(string(content[start:lineEnd]))
}
