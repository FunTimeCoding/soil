package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/call_format"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
)

func findCallFormatEdits(
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

			ast.Inspect(
				file,
				func(n ast.Node) bool {
					call, okay := n.(*ast.CallExpr)

					if !okay {
						return true
					}

					if seen[call.Lparen] {
						return true
					}

					if !call_format.IsViolation(p.Fset, call) {
						return true
					}

					seen[call.Lparen] = true
					line := p.Fset.Position(call.Lparen).Line
					r.AddConcern(
						concern.NewFile(
							"call_format",
							fmt.Sprintf("formatted call (line %d)", line),
							name,
							true,
						),
					)
					fixes := call_format.BuildFixes(
						p.Fset,
						call,
						source,
					)

					for _, f := range fixes {
						result = append(
							result,
							edit{
								position: f.Position,
								end:      f.End,
								newText:  f.NewText,
							},
						)
					}

					return true
				},
			)
		}
	}

	return result
}
