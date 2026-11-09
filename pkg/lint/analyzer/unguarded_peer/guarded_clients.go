package unguarded_peer

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func guardedClients(p *packages.Package) map[types.Object]bool {
	result := make(map[types.Object]bool)

	for _, file := range p.Syntax {
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				switch v := n.(type) {
				case *ast.StructType:
					addClients(p, v.Fields, result)
				case *ast.FuncType:
					addClients(p, v.Params, result)
				}

				return true
			},
		)
	}

	return result
}
