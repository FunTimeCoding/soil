package element_format

import "github.com/dave/dst"

func hasLineComments(items []dst.Expr) bool {
	for _, item := range items {
		if len(item.Decorations().End.All()) > 0 {
			return true
		}
	}

	return false
}
