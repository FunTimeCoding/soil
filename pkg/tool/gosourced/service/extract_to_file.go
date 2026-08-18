package service

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/goast"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/strings/camel"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func (s *Service) ExtractToFile(
	directory string,
	filePath string,
	functionName string,
	dryRun bool,
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
				fmt.Sprintf("would leave an empty file: %s", filePath),
				filePath,
				false,
			),
		)

		return r, nil
	}

	targetPath := filepath.Join(
		filepath.Dir(fullPath),
		fmt.Sprintf("%s.go", camel.ToSnake(functionName)),
	)

	if _, e := os.Stat(targetPath); e == nil {
		r.AddConcern(
			concern.NewFile(
				"validation",
				fmt.Sprintf("%s already exists", filepath.Base(targetPath)),
				filePath,
				false,
			),
		)

		return r, nil
	}

	dec := decorator.NewDecoratorWithImports(
		fileSet,
		constant.StandalonePath,
		goast.New(),
	)
	source, e := dec.DecorateFile(file)

	if e != nil {
		return nil, e
	}

	moved, _ := dec.Dst.Nodes[declaration].(*dst.FuncDecl)

	if moved == nil {
		return nil, fmt.Errorf("no decorated declaration for %s", functionName)
	}

	for i, d := range source.Decls {
		if d == moved {
			source.Decls = append(source.Decls[:i], source.Decls[i+1:]...)

			break
		}
	}

	file.Decls = append(file.Decls[:index], file.Decls[index+1:]...)
	moved.Decs.Before = dst.EmptyLine
	target := &dst.File{
		Name:  dst.NewIdent(file.Name.Name),
		Decls: []dst.Decl{moved},
	}

	if e := restoreExtracted(source, fullPath, dryRun); e != nil {
		return nil, e
	}

	if e := restoreExtracted(target, targetPath, dryRun); e != nil {
		return nil, e
	}

	r.AddConcern(
		concern.NewFile(
			"extracted",
			fmt.Sprintf("%s → %s", functionName, filepath.Base(targetPath)),
			filePath,
			true,
		),
	)

	if countFunctions(file) == 1 {
		name := remainingFunctionName(file)
		renamePath := filepath.Join(
			filepath.Dir(fullPath),
			fmt.Sprintf("%s.go", camel.ToSnake(name)),
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

		if !dryRun {
			e = os.Rename(fullPath, renamePath)

			if e != nil {
				return nil, e
			}
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

	if dryRun {
		r.MarkPlanned()
	}

	return r, nil
}
