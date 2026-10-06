package service

import "github.com/dave/dst"

func removeDeclaredParameters(
	f *dst.FuncDecl,
	indices map[int]bool,
) {
	var fields []*dst.Field
	index := 0

	for _, field := range f.Type.Params.List {
		if len(field.Names) == 0 {
			if !indices[index] {
				fields = append(fields, field)
			}

			index++

			continue
		}

		var names []*dst.Ident

		for _, n := range field.Names {
			if !indices[index] {
				names = append(names, n)
			}

			index++
		}

		if len(names) > 0 {
			field.Names = names
			fields = append(fields, field)
		}
	}

	f.Type.Params.List = fields
}
