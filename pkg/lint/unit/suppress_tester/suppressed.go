package suppress_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"go/ast"
)

func (o *Tester) Suppressed(
	n ast.Node,
	key string,
) bool {
	return suppress.IsSuppressed(o.fileSet, []*ast.File{o.file}, n, key)
}
