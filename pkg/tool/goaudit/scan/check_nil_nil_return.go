package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"go/ast"
	"path/filepath"
)

func (s *Service) checkNilNilReturn(
	f *ast.File,
	filePath string,
) {
	ast.Inspect(
		f,
		func(n ast.Node) bool {
			r, okay := n.(*ast.ReturnStmt)

			if !okay || len(r.Results) != 2 {
				return true
			}

			first, firstOkay := r.Results[0].(*ast.Ident)
			second, secondOkay := r.Results[1].(*ast.Ident)

			if firstOkay && secondOkay && first.Name == "nil" && second.Name == "nil" {
				s.Concerns = append(
					s.Concerns,
					concern.NewPackage(
						NilNilReturnKey,
						fmt.Sprintf(
							"%s: return nil, nil (silent empty response)",
							filepath.Base(filePath),
						),
						filepath.Dir(filepath.Dir(filePath)),
					),
				)
			}

			return true
		},
	)
}
