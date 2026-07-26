package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strconv"
	"strings"
)

func retargetImports(
	r *output.Results,
	decorations *decoration.Set,
	set *token.FileSet,
	all []*packages.Package,
	packagePath string,
	targetPackagePath string,
) error {
	prefix := join.Empty(packagePath, "/")

	for _, loaded := range all {
		for _, file := range loaded.Syntax {
			touched := false

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

				touched = true
				moved := join.Empty(
					targetPackagePath,
					strings.TrimPrefix(importPath, packagePath),
				)
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

			if !touched {
				continue
			}

			decorated, e := decorations.DecorateFile(set, loaded, file)

			if e != nil {
				return e
			}

			decoration.SwapPaths(
				decorated,
				packagePath,
				targetPackagePath,
			)
		}
	}

	return nil
}
