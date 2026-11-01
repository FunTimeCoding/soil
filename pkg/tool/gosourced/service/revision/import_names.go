package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"strconv"
)

func importNames(
	f *ast.File,
	clauses map[string]string,
) map[string]string {
	result := map[string]string{}

	for _, spec := range f.Imports {
		p, e := strconv.Unquote(spec.Path.Value)

		if e != nil {
			continue
		}

		if spec.Name == nil {
			name, inModule := clauses[p]

			if !inModule {
				name = guessName(p)
			}

			result[name] = p

			continue
		}

		if spec.Name.Name != constant.BlankName &&
			spec.Name.Name != constant.DotName {
			result[spec.Name.Name] = p
		}
	}

	return result
}
