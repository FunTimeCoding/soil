package revision

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"go/ast"
	"go/printer"
	"go/token"
)

func render(
	s *token.FileSet,
	n ast.Node,
) string {
	var b bytes.Buffer
	errors.PanicOnError(printer.Fprint(&b, s, n))

	return b.String()
}
