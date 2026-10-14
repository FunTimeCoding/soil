package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func transfers(
	p *packages.Package,
	body *ast.BlockStmt,
) []resource_transfer.Transfer {
	var result []resource_transfer.Transfer
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.CompositeLit:
				result = append(result, literalTransfers(p, v)...)
			case *ast.AssignStmt:
				result = append(result, assignmentTransfers(p, v)...)
			}

			return true
		},
	)

	return result
}
