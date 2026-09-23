package expected_first

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"strings"
)

func expectedIndex(fields *ast.FieldList) int {
	flat := 0

	for _, f := range fields.List {
		for _, n := range f.Names {
			if strings.HasPrefix(n.Name, constant.ExpectedParameterPrefix) {
				return flat
			}

			flat++
		}
	}

	return -1
}
