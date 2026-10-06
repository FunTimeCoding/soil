package service

import (
	"fmt"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"go/parser"
	"go/token"
	"path/filepath"
)

func removeImportThrough(
	directory string,
	filePath string,
	importPath string,
	out *sink.Sink,
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

	if !decoration.RemoveImport(file, importPath) {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernValidation,
				fmt.Sprintf("import %s not found in %s", importPath, filePath),
				filePath,
				false,
			),
		)

		return r, nil
	}

	e = decoration.WriteFile(file, fullPath, out)

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
