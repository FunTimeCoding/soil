package call_format

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/element_format"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/token"
)

func IsViolation(
	fileSet *token.FileSet,
	call *ast.CallExpr,
) bool {
	if len(call.Args) < 2 {
		return false
	}

	openLine := fileSet.Position(call.Lparen).Line
	closeLine := fileSet.Position(call.Rparen).Line

	if openLine == closeLine {
		return lineLength(fileSet, call) > constant.MaxLineLength
	}

	return element_format.IsMultiLineViolation(
		fileSet,
		element_format.FromCall(call),
	)
}
