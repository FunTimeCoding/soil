package tester_function

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

	var found []*ast.FuncDecl

	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay ||
				f.Recv != nil ||
				!f.Name.IsExported() ||
				f.Name.Name == "New" ||
				!takesTestingFirst(f) {
				continue
			}

			found = append(found, f)
		}
	}

	if len(found) < constant.TesterFunctionMinimum {
		return
	}

	for _, f := range found {
		results.AddConcern(
			concern.NewPosition(
				"tester_function",
				fmt.Sprintf(
					"%s is one of %d functions taking *testing.T first; hand out an object built by New(t) and make it a method",
					f.Name.Name,
					len(found),
				),
				p.Fset.Position(f.Pos()),
			),
		)
	}
}
