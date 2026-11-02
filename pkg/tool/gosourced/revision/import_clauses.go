package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"go/token"
	"strings"
)

func importClauses(
	s *token.FileSet,
	parsed []*ast.File,
	module string,
) map[string]string {
	result := map[string]string{}

	for _, f := range parsed {
		if strings.HasSuffix(f.Name.Name, constant.TestClauseSuffix) {
			continue
		}

		result[importPath(module, packageDirectory(s, f))] = f.Name.Name
	}

	return result
}
