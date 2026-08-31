package match

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"go/ast"
	"go/printer"
	"strings"
)

func (m *Matcher) render(expression ast.Expr) string {
	var b strings.Builder
	errors.PanicOnError(printer.Fprint(&b, m.set, expression))

	return b.String()
}
