package symbol

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func receiverTypeName(
	p *packages.Package,
	method *ast.FuncDecl,
) *types.TypeName {
	if method.Recv == nil || len(method.Recv.List) == 0 {
		return nil
	}

	f, okay := p.TypesInfo.Defs[method.Name].(*types.Func)

	if !okay {
		return nil
	}

	t := f.Type().(*types.Signature).Recv().Type()

	if pointer, isPointer := t.(*types.Pointer); isPointer {
		t = pointer.Elem()
	}

	named, okay := t.(*types.Named)

	if !okay {
		return nil
	}

	return named.Obj()
}
