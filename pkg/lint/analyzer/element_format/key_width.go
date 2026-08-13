package element_format

import "github.com/dave/dst"

func KeyWidth(key dst.Expr) int {
	switch k := key.(type) {
	case *dst.Ident:
		return len(k.Name)
	case *dst.BasicLit:
		return len(k.Value)
	}

	return 0
}
