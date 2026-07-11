package service

import (
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/ast"
	"path"
	"strconv"
)

func removeImportIfUnused(
	file *ast.File,
	importPath string,
) {
	var spec *ast.ImportSpec
	local := path.Base(importPath)

	for _, s := range file.Imports {
		p, e := strconv.Unquote(s.Path.Value)

		if e != nil || p != importPath {
			continue
		}

		spec = s

		if s.Name != nil {
			local = s.Name.Name
		}

		break
	}

	if spec == nil {
		return
	}

	used := false
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			s, okay := n.(*ast.SelectorExpr)

			if !okay {
				return true
			}

			x, okay := s.X.(*ast.Ident)

			if okay && x.Name == local {
				used = true
			}

			return true
		},
	)

	if !used {
		imports.Remove(file, spec)
	}
}
