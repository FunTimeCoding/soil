package service

import (
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/ast"
)

func removeUnusedImports(
	file *ast.File,
	extracted []*ast.ImportSpec,
	remaining map[string]bool,
) {
	for _, spec := range extracted {
		if !remaining[spec.Path.Value] {
			imports.Remove(file, spec)
		}
	}
}
