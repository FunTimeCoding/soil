package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strconv"
	"strings"
)

func rewriteImportPaths(
	r *output.Results,
	set *token.FileSet,
	all []*packages.Package,
	packagePath string,
	targetPackagePath string,
	modified map[string]*ast.File,
) {
	prefix := fmt.Sprintf("%s/", packagePath)

	for _, loaded := range all {
		for _, file := range loaded.Syntax {
			filename := set.Position(file.Pos()).Filename

			for _, spec := range file.Imports {
				importPath, e := strconv.Unquote(spec.Path.Value)

				if e != nil {
					continue
				}

				match := importPath == packagePath ||
					strings.HasPrefix(importPath, prefix)

				if !match {
					continue
				}

				moved := fmt.Sprintf(
					"%s%s",
					targetPackagePath,
					strings.TrimPrefix(importPath, packagePath),
				)
				spec.Path.Value = strconv.Quote(moved)
				modified[filename] = file
				position := set.Position(spec.Path.Pos())
				r.AddConcern(
					concern.NewLine(
						"rewritten",
						fmt.Sprintf("%s → %s", importPath, moved),
						position.Filename,
						position.Line,
						"",
						true,
					),
				)
			}
		}
	}
}
