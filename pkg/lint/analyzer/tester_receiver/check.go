package tester_receiver

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	if !strings.HasSuffix(p.Name, constant.TesterSuffix) {
		return
	}

	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || f.Recv == nil || !takesTesting(f) {
				continue
			}

			results.AddConcern(
				concern.NewPosition(
					"tester_receiver",
					fmt.Sprintf(
						"tester method %s takes *testing.T; hold it on the receiver",
						f.Name.Name,
					),
					p.Fset.Position(f.Pos()),
				),
			)
		}
	}
}
