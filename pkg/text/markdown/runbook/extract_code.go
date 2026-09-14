package runbook

import (
	"github.com/yuin/goldmark/v2/ast"
	"strings"
)

func extractCode(
	source *[]byte,
	c *ast.CodeBlock,
) string {
	return strings.TrimSpace(c.Value.Str(*source))
}
