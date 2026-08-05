package stray_variable

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

			if !okay || g.Tok != token.VAR {
				continue
			}

			for _, s := range g.Specs {
				v := s.(*ast.ValueSpec)

				if isEmbedded(g, v) {
					continue
				}

				for i, n := range v.Names {
					if isSynchronization(p, n) {
						continue
					}

					if isOnceConstructed(p, v, i) {
						continue
					}

					if isMetricRegistration(p, v, i) {
						continue
					}

					if isBuildMetadata(p, n) {
						continue
					}

					results.AddConcern(
						concern.NewFile(
							"stray_variable",
							fmt.Sprintf(
								"package-level variable %s outside constant/",
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
