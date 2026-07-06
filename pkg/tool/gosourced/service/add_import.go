package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/parser"
	"go/token"
	"path/filepath"
)

func (s *Service) AddImport(
	directory string,
	filePath string,
	importPath string,
	alias string,
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

	imports.Add(file, importPath, alias)
	e = writeFile(fileSet, file, fullPath)

	if e != nil {
		return nil, e
	}

	message := fmt.Sprintf("added %s", importPath)

	if alias != "" {
		message = fmt.Sprintf("added %s as %s", importPath, alias)
	}

	r.AddConcern(concern.NewFile("import", message, filePath, true))

	return r, nil
}
