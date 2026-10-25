package nil_dependency

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

func checkCall(
	p *packages.Package,
	results *output.Results,
	call *ast.CallExpr,
) {
	name := calleeName(call)

	if !strings.HasPrefix(name, "New") {
		return
	}

	s, okay := p.TypesInfo.TypeOf(call.Fun).(*types.Signature)

	if !okay {
		return
	}

	for i, argument := range call.Args {
		if !p.TypesInfo.Types[argument].IsNil() {
			continue
		}

		dependency, okay := dependencyFace(parameterType(s, call, i))

		if !okay {
			continue
		}

		if suppress.IsSuppressed(
			p.Fset,
			p.Syntax,
			argument.Pos(),
			"nil_dependency",
		) {
			continue
		}

		results.AddConcern(
			concern.NewPosition(
				"nil_dependency",
				fmt.Sprintf(
					"nil passed to %s as %s - pass a mock or a real one",
					types.ExprString(call.Fun),
					dependency,
				),
				p.Fset.Position(argument.Pos()),
			),
		)
	}
}
