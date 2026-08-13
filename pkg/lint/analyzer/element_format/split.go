package element_format

import "github.com/dave/dst"

func Split(items []dst.Expr) {
	for _, item := range items {
		item.Decorations().Before = dst.NewLine
		item.Decorations().After = dst.NewLine
	}
}
