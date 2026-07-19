package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"path"
)

func chooseImportName(
	file *ast.File,
	importPath string,
	packageName string,
) *importName {
	names := importLocalNames(file)

	for local, p := range names {
		if p == importPath {
			return &importName{local: local, imported: true}
		}
	}

	subsystem := path.Base(path.Dir(importPath))
	candidates := []string{
		packageName,
		subsystem,
		join.Empty(subsystem, FlipName(packageName)),
	}

	for _, c := range candidates {
		if _, taken := names[c]; !taken {
			result := &importName{local: c}

			if c != packageName {
				result.alias = c
			}

			return result
		}
	}

	return nil
}
