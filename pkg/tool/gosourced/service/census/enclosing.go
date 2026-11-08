package census

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"go/token"
)

func enclosing(
	f *ast.File,
	at token.Pos,
) string {
	for _, d := range f.Decls {
		if at < d.Pos() || at > d.End() {
			continue
		}

		x, okay := d.(*ast.FuncDecl)

		if !okay {
			return constant.PackageScope
		}

		if r := receiverType(x); r != "" {
			return join.Empty(r, constant.MemberSeparator, x.Name.Name)
		}

		return x.Name.Name
	}

	return constant.PackageScope
}
