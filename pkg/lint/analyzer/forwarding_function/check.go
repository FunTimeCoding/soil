package forwarding_function

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay {
				continue
			}

			call := forwardedCall(f)

			if call == nil || call.Ellipsis.IsValid() {
				continue
			}

			if !plainCallee(call.Fun) {
				continue
			}

			if !forwardsParameters(f, call) {
				continue
			}

			name := calledName(call.Fun)

			if name == "" || name == f.Name.Name {
				continue
			}

			results.AddConcern(
				concern.NewPosition(
					"forwarding_function",
					fmt.Sprintf(
						"%s only forwards to %s; call it directly",
						f.Name.Name,
						name,
					),
					p.Fset.Position(f.Pos()),
				),
			)
		}
	}
}
