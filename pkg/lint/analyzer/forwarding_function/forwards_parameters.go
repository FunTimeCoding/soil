package forwarding_function

import "go/ast"

func forwardsParameters(
	f *ast.FuncDecl,
	call *ast.CallExpr,
) bool {
	names := parameterNames(f)

	if len(names) == 0 || len(call.Args) != len(names) {
		return false
	}

	for i, argument := range call.Args {
		identifier, okay := argument.(*ast.Ident)

		if !okay || identifier.Name != names[i] {
			return false
		}
	}

	return true
}
