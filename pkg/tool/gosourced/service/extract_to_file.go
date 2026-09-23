package service

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/goast"
	"github.com/dave/dst/decorator/resolver/gopackages"
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/strings/camel"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
	"strings"
)

func (s *Service) ExtractToFile(
	directory string,
	filePath string,
	symbolName string,
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

	var declaration ast.Decl
	functionDeclaration, index := findFunctionDeclaration(file, symbolName)

	if functionDeclaration != nil {
		declaration = functionDeclaration
	}

	if declaration == nil {
		typeDeclaration, typeIndex, grouped := findTypeDeclaration(
			file,
			symbolName,
		)

		if grouped {
			r.AddConcern(
				concern.NewFile(
					"validation",
					fmt.Sprintf("%s is declared in a type group", symbolName),
					filePath,
					false,
				),
			)

			return r, nil
		}

		if typeDeclaration != nil {
			declaration = typeDeclaration
			index = typeIndex
		}
	}

	if declaration == nil {
		r.AddConcern(
			concern.NewFile(
				"validation",
				fmt.Sprintf("symbol %s not found in %s", symbolName, filePath),
				filePath,
				false,
			),
		)

		return r, nil
	}

	if !hasCompanionDeclaration(file, index) {
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
		fmt.Sprintf("%s.go", camel.ToSnake(symbolName)),
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
		goast.WithResolver(
			gopackages.WithConfig(
				filepath.Dir(fullPath),
				packages.Config{BuildFlags: resolve.BuildFlags(directory)},
			),
		),
	)
	source, e := dec.DecorateFile(file)

	if e != nil {
		return nil, e
	}

	moved, _ := dec.Dst.Nodes[declaration].(dst.Decl)

	if moved == nil {
		return nil, not_found.Format(
			"no decorated declaration for %s",
			symbolName,
		)
	}

	for i, d := range source.Decls {
		if d == moved {
			source.Decls = append(source.Decls[:i], source.Decls[i+1:]...)

			break
		}
	}

	file.Decls = append(file.Decls[:index], file.Decls[index+1:]...)
	moved.Decorations().Before = dst.EmptyLine
	target := &dst.File{
		Name:  dst.NewIdent(file.Name.Name),
		Decls: []dst.Decl{moved},
	}

	for _, line := range buildConstraints(file) {
		target.Decs.Start.Append(line, "\n")
	}

	if e := restoreExtracted(directory, source, fullPath, dryRun); e != nil {
		return nil, e
	}

	if e := restoreExtracted(directory, target, targetPath, dryRun); e != nil {
		return nil, e
	}

	r.AddConcern(
		concern.NewFile(
			"extracted",
			fmt.Sprintf("%s → %s", symbolName, filepath.Base(targetPath)),
			filePath,
			true,
		),
	)

	if countIdentities(file) == 1 {
		if strings.HasSuffix(fullPath, libraryConstant.TestSuffix) {
			r.AddConcern(
				concern.NewFile(
					"extracted",
					fmt.Sprintf(
						"%s was not renamed: test file naming is not one-identity-per-file; rename it by hand if %s no longer fits",
						filepath.Base(fullPath),
						filepath.Base(fullPath),
					),
					filePath,
					true,
				),
			)

			if dryRun {
				r.MarkPlanned()
			}

			return r, nil
		}

		name := remainingIdentityName(file)
		renamePath := filepath.Join(
			filepath.Dir(fullPath),
			fmt.Sprintf("%s.go", camel.ToSnake(name)),
		)

		if renamePath == fullPath {
			if dryRun {
				r.MarkPlanned()
			}

			return r, nil
		}

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
