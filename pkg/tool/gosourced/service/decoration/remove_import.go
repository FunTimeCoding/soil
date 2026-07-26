package decoration

import (
	"fmt"
	"github.com/dave/dst"
	"go/token"
)

func RemoveImport(
	file *dst.File,
	importPath string,
) bool {
	quoted := fmt.Sprintf("%q", importPath)

	for i, d := range file.Decls {
		declaration, okay := d.(*dst.GenDecl)

		if !okay || declaration.Tok != token.IMPORT {
			continue
		}

		for j, s := range declaration.Specs {
			candidate := s.(*dst.ImportSpec)

			if candidate.Path.Value != quoted {
				continue
			}

			declaration.Specs = append(
				declaration.Specs[:j],
				declaration.Specs[j+1:]...,
			)

			if len(declaration.Specs) == 0 {
				file.Decls = append(
					file.Decls[:i],
					file.Decls[i+1:]...,
				)
			}

			return true
		}
	}

	return false
}
