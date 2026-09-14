package runbook

import (
	"github.com/yuin/goldmark/v2/ast"
	"strings"
)

func extractText(
	source *[]byte,
	n ast.Node,
) string {
	var b strings.Builder

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() == ast.KindText {
			b.WriteString(c.(*ast.Text).Value.Value(*source))
		}
	}

	return b.String()
}
