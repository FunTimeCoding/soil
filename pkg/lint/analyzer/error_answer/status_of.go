package error_answer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strconv"
)

func statusOf(
	p *packages.Package,
	e ast.Expr,
) int {
	var t types.Type

	if x, okay := e.(*ast.CompositeLit); okay {
		t = p.TypesInfo.TypeOf(x)
	}

	if x, okay := e.(*ast.CallExpr); okay {
		if v, known := p.TypesInfo.Types[x.Fun]; known && v.IsType() {
			t = v.Type
		}
	}

	named, okay := t.(*types.Named)

	if !okay {
		return 0
	}

	m := constant.ResponseStatus.FindStringSubmatch(named.Obj().Name())

	if m == nil {
		return 0
	}

	status, f := strconv.Atoi(m[1])

	if f != nil {
		return 0
	}

	return status
}
