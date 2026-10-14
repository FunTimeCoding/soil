package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func returnedTransfers(
	p *packages.Package,
	f *ast.FuncDecl,
) []resource_transfer.Transfer {
	var result []resource_transfer.Transfer
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.FuncLit:
				return false
			case *ast.ReturnStmt:
				for _, e := range v.Results {
					result = append(result, valueTransfers(p, f.Body, e)...)
				}
			}

			return true
		},
	)

	return result
}
