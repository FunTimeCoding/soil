package service

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"os"
)

func writeExtractedFile(
	packageName string,
	needed []*ast.ImportSpec,
	declarations []ast.Decl,
	path string,
) error {
	fileSet := token.NewFileSet()
	file := &ast.File{
		Name: ast.NewIdent(packageName),
	}

	if len(needed) > 0 {
		specs := make([]ast.Spec, len(needed))

		for i, spec := range needed {
			specs[i] = spec
		}

		file.Decls = append(
			file.Decls,
			&ast.GenDecl{
				Tok:   token.IMPORT,
				Specs: specs,
			},
		)
	}

	file.Decls = append(file.Decls, declarations...)
	var buffer bytes.Buffer
	e := format.Node(&buffer, fileSet, file)

	if e != nil {
		return e
	}

	return os.WriteFile(path, buffer.Bytes(), 0644)
}
