package struct_literal

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

func ownedType(
	p *packages.Package,
	sel *ast.SelectorExpr,
	module string,
) *types.TypeName {
	o, okay := p.TypesInfo.Uses[sel.Sel]

	if !okay {
		return nil
	}

	typeName, okay := o.(*types.TypeName)

	if !okay {
		return nil
	}

	_, isStruct := typeName.Type().Underlying().(*types.Struct)

	if !isStruct {
		return nil
	}

	a := typeName.Pkg()

	if a == nil {
		return nil
	}

	path := a.Path()

	if path != module && !strings.HasPrefix(
		path,
		fmt.Sprintf("%s/", module),
	) {
		return nil
	}

	filename := p.Fset.Position(typeName.Pos()).Filename

	for _, file := range p.Syntax {
		if p.Fset.File(file.Pos()).Name() == filename && ast.IsGenerated(file) {
			return nil
		}
	}

	if lint.IsGeneratedFile(filename) {
		return nil
	}

	return typeName
}
