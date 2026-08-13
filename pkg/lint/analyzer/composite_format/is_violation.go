package composite_format

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/element_format"
	"go/ast"
	"go/token"
)

func IsViolation(
	fileSet *token.FileSet,
	lit *ast.CompositeLit,
) bool {
	if len(lit.Elts) < 2 {
		return false
	}

	openLine := fileSet.Position(lit.Lbrace).Line
	closeLine := fileSet.Position(lit.Rbrace).Line

	if openLine == closeLine {
		return false
	}

	return element_format.IsMultiLineViolation(
		fileSet,
		element_format.FromLiteral(lit),
	)
}
