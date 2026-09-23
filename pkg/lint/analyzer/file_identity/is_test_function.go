package file_identity

import (
	"go/ast"
	"strings"
)

func isTestFunction(f *ast.FuncDecl) bool {
	if f.Recv != nil {
		return false
	}

	for _, prefix := range []string{
		"Test",
		"Benchmark",
		"Fuzz",
		"Example",
	} {
		if strings.HasPrefix(f.Name.Name, prefix) {
			return true
		}
	}

	return false
}
