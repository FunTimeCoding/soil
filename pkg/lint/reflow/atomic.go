package reflow

import "github.com/yuin/goldmark/v2/ast"

func atomic(
	source []byte,
	document ast.Node,
) []bool {
	result := literal(source, document)

	for i, v := range linked(source, document) {
		result[i] = result[i] || v
	}

	return result
}
