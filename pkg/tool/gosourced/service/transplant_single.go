package service

import "github.com/dave/dst"

func transplantSingle(
	declaration *dst.GenDecl,
	spec dst.Spec,
	single bool,
) dst.Decl {
	if single {
		declaration.Decs.Before = dst.EmptyLine

		return declaration
	}

	fresh := &dst.GenDecl{
		Tok:   declaration.Tok,
		Specs: []dst.Spec{spec},
	}
	fresh.Decs.Before = dst.EmptyLine
	fresh.Decs.Start.Append(spec.Decorations().Start.All()...)
	spec.Decorations().Start.Clear()
	spec.Decorations().Before = dst.None

	return fresh
}
