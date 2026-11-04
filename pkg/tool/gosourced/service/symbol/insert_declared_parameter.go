package symbol

import "github.com/dave/dst"

func insertDeclaredParameter(
	f *dst.FuncDecl,
	position int,
	name string,
	t dst.Expr,
) {
	var fields []*dst.Field
	index := 0
	added := false
	multiline := false
	add := func() {
		fields = append(
			fields,
			&dst.Field{Names: []*dst.Ident{dst.NewIdent(name)}, Type: t},
		)
		added = true
	}

	for _, field := range f.Type.Params.List {
		if field.Decs.Before == dst.NewLine {
			multiline = true
		}

		count := len(field.Names)

		if !added && position == index {
			add()
		}

		if !added && position > index && position < index+count {
			split := position - index
			fields = append(
				fields,
				&dst.Field{
					Names: field.Names[:split],
					Type:  dst.Clone(field.Type).(dst.Expr),
				},
			)
			add()
			field.Names = field.Names[split:]
		}

		fields = append(fields, field)
		index += count
	}

	if !added {
		add()
	}

	if multiline {
		for i, field := range fields {
			field.Decs.Before = dst.NewLine
			field.Decs.After = dst.None

			if i == len(fields)-1 {
				field.Decs.After = dst.NewLine
			}
		}
	}

	f.Type.Params.List = fields
}
