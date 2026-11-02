package forwarding_function

import "go/ast"

func forwards(f *ast.FuncDecl) *ast.CallExpr {
	call := forwardedCall(f)

	if call == nil || call.Ellipsis.IsValid() {
		return nil
	}

	if !plainCallee(call.Fun) || !forwardsParameters(f, call) {
		return nil
	}

	return call
}
