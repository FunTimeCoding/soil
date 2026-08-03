package expected_first

import (
	"go/ast"
	"strings"
)

// expectedIndex returns the flat parameter position of the first
// expected-named parameter, or -1 when none exists.
func expectedIndex(fields *ast.FieldList) int {
	flat := 0

	for _, f := range fields.List {
		for _, n := range f.Names {
			if strings.HasPrefix(n.Name, "expected") {
				return flat
			}

			flat++
		}
	}

	return -1
}
