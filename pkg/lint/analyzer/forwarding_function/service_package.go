package forwarding_function

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"slices"
	"strings"
)

func servicePackage(
	p *packages.Package,
	receiver ast.Expr,
) string {
	t := p.TypesInfo.TypeOf(receiver)

	if t == nil {
		return ""
	}

	if pointer, okay := t.(*types.Pointer); okay {
		t = pointer.Elem()
	}

	named, okay := t.(*types.Named)

	if !okay || named.Obj().Pkg() == nil {
		return ""
	}

	owner := named.Obj().Pkg()

	if !slices.Contains(
		strings.Split(owner.Path(), "/"),
		constant.ServiceDirectory,
	) {
		return ""
	}

	return owner.Name()
}
