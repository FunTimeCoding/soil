package call_format

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/token"
)

// BuildFixes returns the edits needed to fix a call_format violation.
func BuildFixes(
	fileSet *token.FileSet,
	call *ast.CallExpr,
	source []byte,
) []Fix {
	indent := callIndent(fileSet, call, source)
	argIndent := join.Empty(indent, "\t")
	openLine := fileSet.Position(call.Lparen).Line
	closeLine := fileSet.Position(call.Rparen).Line

	if openLine == closeLine {
		return fixSingleLine(call, argIndent, indent)
	}

	return fixMultiLine(fileSet, call, argIndent)
}
