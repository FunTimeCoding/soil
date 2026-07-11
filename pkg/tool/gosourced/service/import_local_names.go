package service

import (
	"go/ast"
	"path"
	"strconv"
)

func importLocalNames(file *ast.File) map[string]string {
	result := make(map[string]string)

	for _, spec := range file.Imports {
		importPath, e := strconv.Unquote(spec.Path.Value)

		if e != nil {
			continue
		}

		local := path.Base(importPath)

		if spec.Name != nil {
			local = spec.Name.Name
		}

		result[local] = importPath
	}

	return result
}
