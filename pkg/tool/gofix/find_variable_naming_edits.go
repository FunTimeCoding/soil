package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/variable_naming"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func findVariableNamingEdits(
	fileSet *token.FileSet,
	all []*packages.Package,
	r *output.Results,
) []edit {
	var result []edit
	seen := make(map[token.Pos]bool)

	for _, p := range all {
		for _, file := range p.Syntax {
			if filepath.Base(p.Fset.File(file.Pos()).Name()) == constant.GeneratedFile {
				continue
			}

			if ast.IsGenerated(file) {
				continue
			}

			ast.Inspect(
				file,
				func(n ast.Node) bool {
					f, okay := n.(*ast.FuncDecl)

					if !okay || f.Body == nil {
						return true
					}

					renames := variable_naming.ComputeRenames(p, f)

					for _, rename := range renames {
						if seen[rename.Object.Pos()] {
							continue
						}

						seen[rename.Object.Pos()] = true
						path := fileSet.File(rename.Object.Pos()).Name()
						r.AddConcern(
							concern.NewFile(
								"variable_naming",
								fmt.Sprintf(
									"renamed %s → %s",
									rename.Object.Name(),
									rename.NewName,
								),
								path,
								true,
							),
						)
						references := resolve.FindAllReferences(
							all,
							rename.Object,
						)

						for _, e := range references {
							result = append(
								result,
								edit{
									position: e.Ident.Pos(),
									end:      e.Ident.End(),
									newText:  rename.NewName,
								},
							)
						}
					}

					return true
				},
			)
		}
	}

	return result
}
