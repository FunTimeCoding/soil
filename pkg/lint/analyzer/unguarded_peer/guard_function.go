package unguarded_peer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func guardFunction(
	p *packages.Package,
	f ast.Expr,
) bool {
	o, okay := callee(p, f).(*types.Func)

	if !okay || o.Pkg() == nil || o.Pkg().Path() != constant.PeerPackage {
		return false
	}

	s := o.Type().(*types.Signature)

	return s.Recv() == nil && takesGuard(s) && takesFunction(s)
}
