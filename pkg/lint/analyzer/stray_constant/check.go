package stray_constant

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
)

// Check flags package-level constants outside constant/ packages
// - the twin of stray_variable for the const keyword. Constants
// live in the domain's constant/ home; sanctions accumulate here
// as the census surfaces legitimate residents.
func Check(
	p *packages.Package,
	results *output.Results,
) {
	// Test-variant package entries recompile the same source
	// files - visiting them would double-report every constant.
	if len(p.Syntax) == 0 || p.ID != p.PkgPath {
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
