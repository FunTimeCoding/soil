package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
)

func (s *Service) RemoveImport(
	directory string,
	filePath string,
	importPath string,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	fullPath := filePath

	if !filepath.IsAbs(fullPath) {
		fullPath = filepath.Join(directory, filePath)
	}

	fileSet := token.NewFileSet()
	file, e := parser.ParseFile(fileSet, fullPath, nil, parser.ParseComments)

	if e != nil {
		return nil, e
	}

	quoted := fmt.Sprintf("%q", importPath)
	var found *ast.ImportSpec

	for _, spec := range file.Imports {
		if spec.Path.Value == quoted {
			found = spec

			break
		}
	}

	if found == nil {
		r.AddConcern(
			concern.NewFile(
				"validation",
				fmt.Sprintf(
					"import %s not found in %s",
					importPath,
					filePath,
				),
				filePath,
				false,
			),
		)

		return r, nil
	}

	imports.Remove(file, found)
	e = writeFile(fileSet, file, fullPath)

	if e != nil {
		return nil, e
	}

	r.AddConcern(
		concern.NewFile(
			"import",
			fmt.Sprintf("removed %s", importPath),
			filePath,
			true,
		),
	)

	return r, nil
}
