package service

import (
	"fmt"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"go/parser"
	"go/token"
	"path/filepath"
)

func (s *Service) AddImport(
	directory string,
	filePath string,
	importPath string,
	alias string,
	dryRun bool,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	fullPath := filePath

	if !filepath.IsAbs(fullPath) {
		fullPath = filepath.Join(directory, filePath)
	}

	fileSet := token.NewFileSet()
	file, e := decorator.ParseFile(fileSet, fullPath, nil, parser.ParseComments)

	if e != nil {
		return nil, e
	}

	decoration.AddImport(file, importPath, alias)
	e = decoration.WriteFile(file, fullPath, dryRun)

	if e != nil {
		return nil, e
	}

	message := fmt.Sprintf("added %s", importPath)

	if alias != "" {
		message = fmt.Sprintf("added %s as %s", importPath, alias)
	}

	r.AddConcern(concern.NewFile("import", message, filePath, true))

	if dryRun {
		r.MarkPlanned()
	}

	return r, nil
}
