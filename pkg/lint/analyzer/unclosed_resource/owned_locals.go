package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func ownedLocals(
	p *packages.Package,
	body *ast.BlockStmt,
	s *Summaries,
) map[*types.Var]bool {
	result := make(map[*types.Var]bool)

	for _, c := range candidates(p, body) {
		if closes(p, body, c.Object, s) || s.isUnowned(p, body, c) {
			continue
		}

		result[c.Object] = true
	}

	return result
}
