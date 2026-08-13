package gofix

import "github.com/dave/dst"

func longestSpecName(declaration *dst.GenDecl) int {
	longest := 0

	for _, spec := range declaration.Specs {
		v, okay := spec.(*dst.ValueSpec)

		if !okay {
			continue
		}

		for _, name := range v.Names {
			if len(name.Name) > longest {
				longest = len(name.Name)
			}
		}
	}

	return longest
}
