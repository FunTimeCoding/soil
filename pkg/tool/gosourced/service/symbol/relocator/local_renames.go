package relocator

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"go/ast"
	"go/types"
)

func localRenames(
	q *relocation.FileQualification,
	shadowing []*types.Var,
	held func(string) bool,
) (map[*ast.Ident]string, *types.Var) {
	result := make(map[*ast.Ident]string)

	for _, v := range shadowing {
		letter := localLetter(q, v, held)

		if letter == "" {
			return nil, v
		}

		for _, i := range localIdents(q, v) {
			result[i] = letter
		}
	}

	return result, nil
}
