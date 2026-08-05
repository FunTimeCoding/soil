package constant_declaration

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	if lint.SkipPackage(p) {
		return
	}

	directory := filepath.Dir(p.Fset.File(p.Syntax[0].Pos()).Name())

	if filepath.Base(directory) != "constant" {
		return
	}

	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		name := p.Fset.File(file.Pos()).Name()

		if strings.HasSuffix(name, constant.TestSuffix) {
			continue
		}

		for _, d := range file.Decls {
			if f, okay := d.(*ast.FuncDecl); okay {
				results.AddConcern(
					concern.NewFile(
						"constant_declaration",
						fmt.Sprintf(
							"func %s inside constant/ - constant packages are behavior-free",
							f.Name.Name,
						),
						name,
						false,
					),
				)

				continue
			}

			g, okay := d.(*ast.GenDecl)

			if !okay || g.Tok != token.TYPE {
				continue
			}

			for _, s := range g.Specs {
				t := s.(*ast.TypeSpec)

				if isEnumShaped(p, t) {
					continue
				}

				results.AddConcern(
					concern.NewFile(
						"constant_declaration",
						fmt.Sprintf(
							"type %s inside constant/ is not enum-shaped - record types live in types/ leaves",
							t.Name.Name,
						),
						name,
						false,
					),
				)
			}
		}
	}
}
