package unit

import (
	"github.com/funtimecoding/soil/pkg/crap/complexity"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func measure(
	t *testing.T,
	body string,
) int {
	t.Helper()
	file, e := parser.ParseFile(
		token.NewFileSet(),
		"f.go",
		join.Empty("package p\n\n", body),
		0,
	)
	errors.PanicOnError(e)

	return complexity.Complexity(file.Decls[0].(*ast.FuncDecl))
}
