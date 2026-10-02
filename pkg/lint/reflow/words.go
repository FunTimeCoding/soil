package reflow

import "github.com/yuin/goldmark/v2/ast"

func words(
	source []byte,
	document ast.Node,
) []string {
	return units(source, literal(source, document), 0, len(source))
}
