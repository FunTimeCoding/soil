package service

import "github.com/dave/dst"

func callReferences(call *dst.CallExpr, name string) bool {
	switch fun := call.Fun.(type) {
	case *dst.Ident:
		return fun.Name == name
	case *dst.SelectorExpr:
		return fun.Sel.Name == name
	}

	return false
}
