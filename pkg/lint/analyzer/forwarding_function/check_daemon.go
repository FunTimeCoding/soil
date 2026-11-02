package forwarding_function

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func CheckDaemon(
	p *packages.Package,
	results *output.Results,
) {
	if p.Name != constant.DaemonPackage {
		return
	}

	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || f.Recv == nil {
				continue
			}

			call := forwards(f)

			if call == nil {
				continue
			}

			s, okay := call.Fun.(*ast.SelectorExpr)

			if !okay || s.Sel.Name != f.Name.Name {
				continue
			}

			service := servicePackage(p, s.X)

			if service == "" {
				continue
			}

			target := fmt.Sprintf("the %s service", service)

			if service == constant.ServiceDirectory {
				target = "the service"
			}

			results.AddConcern(
				concern.NewPosition(
					"forwarding_function",
					fmt.Sprintf(
						"%s only forwards to %s - hand the surface the service",
						f.Name.Name,
						target,
					),
					p.Fset.Position(f.Pos()),
				),
			)
		}
	}
}
