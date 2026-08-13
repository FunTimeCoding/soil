package element_format

import "github.com/dave/dst"

func Collapse(items []dst.Expr) {
	for _, item := range items {
		item.Decorations().Before = dst.None
		item.Decorations().After = dst.None
	}
}
