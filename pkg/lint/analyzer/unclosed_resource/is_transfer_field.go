package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func isTransferField(
	p *packages.Package,
	e ast.Expr,
	t resource_transfer.Transfer,
) bool {
	s, okay := ast.Unparen(e).(*ast.SelectorExpr)

	if !okay || s.Sel.Name != t.Field.Name() {
		return false
	}

	holder := selectedHolder(p, s)

	return holder != nil && holder.Origin() == t.Holder.Origin()
}
