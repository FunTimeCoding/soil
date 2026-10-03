package gofix

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"slices"
)

func formatFiles(all []*packages.Package) []string {
	var result []string

	for _, p := range all {
		for _, file := range p.Syntax {
			name := p.Fset.File(file.Pos()).Name()

			if filepath.Base(name) == constant.GeneratedFile ||
				ast.IsGenerated(file) ||
				!filepath.IsAbs(name) ||
				!fileExists(name) ||
				slices.Contains(result, name) {
				continue
			}

			result = append(result, name)
		}
	}

	return result
}
