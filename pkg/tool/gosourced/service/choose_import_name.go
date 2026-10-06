package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/ast"
	"path"
)

func chooseImportName(
	file *ast.File,
	importPath string,
	packageName string,
) *relocation.ImportName {
	names := importLocalNames(file)

	for local, p := range names {
		if p == importPath {
			return relocation.NewImportName(local, true)
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
			result := relocation.NewImportName(c, false)

			if c != packageName {
				result.Alias = c
			}

			return result
		}
	}

	return nil
}
