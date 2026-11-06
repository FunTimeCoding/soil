package stub_constructor

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	models := gormModels(p)

	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || f.Recv != nil {
				continue
			}

			if f.Name.Name == constant.StubConstructor {
				if f.Type.Params.NumFields() > 0 {
					results.AddConcern(
						concern.NewPosition(
							"stub_constructor",
							"Stub takes no parameters; a constructor building from arguments is New",
							p.Fset.Position(f.Pos()),
						),
					)
				}

				continue
			}

			model := emptyReturn(f)

			if model == "" || !models[model] {
				continue
			}

			results.AddConcern(
				concern.NewPosition(
					"stub_constructor",
					fmt.Sprintf(
						"%s is a gorm model; its empty constructor %s is Stub",
						model,
						f.Name.Name,
					),
					p.Fset.Position(f.Pos()),
				),
			)
		}
	}
}
