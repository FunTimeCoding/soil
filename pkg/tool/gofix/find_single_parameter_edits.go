package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
)

func findSingleParameterEdits(
	all []*packages.Package,
	r *output.Results,
) []edit {
	var result []edit
	sourceCache := make(map[string][]byte)
	seen := make(map[token.Pos]bool)

	for _, p := range all {
		generated := make(map[string]bool)

		for _, file := range p.Syntax {
			name := p.Fset.File(file.Pos()).Name()

			if filepath.Base(name) == constant.GeneratedFile {
				generated[name] = true
			}
		}

		for _, file := range p.Syntax {
			name := p.Fset.File(file.Pos()).Name()

			if generated[name] {
				continue
			}

			if ast.IsGenerated(file) {
				continue
			}

			if !filepath.IsAbs(name) || !fileExists(name) {
				continue
			}

			source, okay := sourceCache[name]

			if !okay {
				var e error
				source, e = os.ReadFile(name)

				if e != nil {
					continue
				}

				sourceCache[name] = source
			}

			for _, declaration := range file.Decls {
				f, okay := declaration.(*ast.FuncDecl)

				if !okay {
					continue
				}

				if seen[f.Type.Params.Opening] {
					continue
				}

				edits := collapseSingleParameter(p.Fset, f, source)

				if edits == nil {
					continue
				}

				seen[f.Type.Params.Opening] = true
				line := p.Fset.Position(f.Pos()).Line
				r.AddConcern(
					concern.NewFile(
						"single_parameter",
						fmt.Sprintf(
							"collapsed single parameter (line %d)",
							line,
						),
						name,
						true,
					),
				)
				result = append(result, edits...)
			}
		}
	}

	return result
}
