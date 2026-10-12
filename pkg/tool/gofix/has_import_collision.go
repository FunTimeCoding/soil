package gofix

import (
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func hasImportCollision(
	file *ast.File,
	p *packages.Package,
	currentPath string,
	declaredName string,
) bool {
	count := 0

	for _, spec := range file.Imports {
		path := strings.Trim(spec.Path.Value, `"`)

		if path == currentPath {
			continue
		}

		if spec.Name != nil && spec.Name.Name != "_" && spec.Name.Name != "." {
			if spec.Name.Name == declaredName {
				return true
			}

			imported := imports.ImportedPackage(p, path)

			if imported != nil && imported.Name() == declaredName {
				count++
			}

			continue
		}

		imported := imports.ImportedPackage(p, path)

		if imported != nil && imported.Name() == declaredName {
			return true
		}
	}

	return count > 0
}
