package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/census/reading"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func readBody(
	p *packages.Package,
	d *ast.FuncDecl,
	file string,
	subject *types.Named,
	module string,
	stores map[string]bool,
) *reading.Reading {
	result := reading.NewReading(d.Name.Name, file)

	for _, f := range d.Type.Params.List {
		if n := named(p.TypesInfo.TypeOf(f.Type)); n != nil && stores[qualified(n)] {
			result.Step = true
		}
	}

	ast.Inspect(
		d.Body,
		func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.SelectorExpr:
				s := p.TypesInfo.Selections[x]

				if s == nil || s.Kind() != types.FieldVal {
					return true
				}

				if r := named(s.Recv()); r != nil && sameType(r, subject) {
					result.Fields[x.Sel.Name] = true
				}
			case *ast.Ident:
				readUse(p, x, subject, module, result)
			}

			return true
		},
	)

	return result
}
