package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func (s *Service) ExtractToFile(
	directory string,
	filePath string,
	functionName string,
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

	declaration, index := findFunctionDeclaration(file, functionName)

	if declaration == nil {
		r.AddConcern(
			concern.NewFile(
				"validation",
				fmt.Sprintf(
					"function %s not found in %s",
					functionName,
					filePath,
				),
				filePath,
				false,
			),
		)

		return r, nil
	}

	if countFunctions(file) == 1 {
		r.AddConcern(
			concern.NewFile(
				"validation",
				fmt.Sprintf(
					"would leave an empty file: %s",
					filePath,
				),
				filePath,
				false,
			),
		)

		return r, nil
	}

	targetPath := filepath.Join(
		filepath.Dir(fullPath),
		fmt.Sprintf("%s.go", toSnakeCase(functionName)),
	)

	if _, e := os.Stat(targetPath); e == nil {
		r.AddConcern(
			concern.NewFile(
				"validation",
				fmt.Sprintf(
					"%s already exists",
					filepath.Base(targetPath),
				),
				filePath,
				false,
			),
		)

		return r, nil
	}

	needed := imports.UsedBy(file, declaration)
	file.Decls = append(file.Decls[:index], file.Decls[index+1:]...)
	remaining := collectRemainingImports(file)
	removeUnusedImports(file, needed, remaining)
	e = writeFile(fileSet, file, fullPath)

	if e != nil {
		return nil, e
	}

	e = writeExtractedFile(
		file.Name.Name,
		needed,
		declaration,
		targetPath,
	)

	if e != nil {
		return nil, e
	}

	r.AddConcern(
		concern.NewFile(
			"extracted",
			fmt.Sprintf(
				"%s → %s",
				functionName,
				filepath.Base(targetPath),
			),
			filePath,
			true,
		),
	)

	if countFunctions(file) == 1 {
		name := remainingFunctionName(file)
		renamePath := filepath.Join(
			filepath.Dir(fullPath),
			fmt.Sprintf("%s.go", toSnakeCase(name)),
		)

		if _, e := os.Stat(renamePath); e == nil {
			r.AddConcern(
				concern.NewFile(
					"validation",
					fmt.Sprintf(
						"cannot rename source: %s already exists",
						filepath.Base(renamePath),
					),
					filePath,
					false,
				),
			)

			return r, nil
		}

		e = os.Rename(fullPath, renamePath)

		if e != nil {
			return nil, e
		}

		r.AddConcern(
			concern.NewFile(
				"renamed",
				fmt.Sprintf(
					"%s → %s",
					filepath.Base(fullPath),
					filepath.Base(renamePath),
				),
				filePath,
				true,
			),
		)
	}

	return r, nil
}
