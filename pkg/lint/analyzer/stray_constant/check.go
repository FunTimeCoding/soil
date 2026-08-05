package stray_constant

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

	if filepath.Base(directory) == "constant" {
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
			g, okay := d.(*ast.GenDecl)

			if !okay || g.Tok != token.CONST {
				continue
			}

			for _, s := range g.Specs {
				v := s.(*ast.ValueSpec)

				for _, n := range v.Names {
					if n.Name == "_" {
						continue
					}

					results.AddConcern(
						concern.NewFile(
							"stray_constant",
							fmt.Sprintf(
								"package-level constant %s outside constant/",
								n.Name,
							),
							name,
							false,
						),
					)
				}
			}
		}
	}
}
