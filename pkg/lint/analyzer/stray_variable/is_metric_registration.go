package stray_variable

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

// Promauto constructors register into the process-global
// prometheus registry at construction - package-level vars are
// the ecosystem's own idiom for them.
func isMetricRegistration(
	p *packages.Package,
	v *ast.ValueSpec,
	index int,
) bool {
	if index >= len(v.Values) {
		return false
	}

	call, okay := v.Values[index].(*ast.CallExpr)

	if !okay {
		return false
	}

	selector, okay := call.Fun.(*ast.SelectorExpr)

	if !okay {
		return false
	}

	x, okay := selector.X.(*ast.Ident)

	if !okay {
		return false
	}

	name, okay := p.TypesInfo.Uses[x].(*types.PkgName)

	return okay &&
		constant.MetricConstructorPackages[name.Imported().Path()]
}
