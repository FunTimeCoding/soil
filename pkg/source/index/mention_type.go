package index

import "go/types"

func mentionType(
	t types.Type,
	self *types.Package,
	into map[string]bool,
	seen map[types.Type]bool,
) {
	if t == nil || seen[t] {
		return
	}

	seen[t] = true
	walk := func(u types.Type) { mentionType(u, self, into, seen) }

	switch v := t.(type) {
	case *types.Alias:
		if p := v.Obj().Pkg(); p != nil && p != self {
			into[p.Path()] = true
		}

		walk(types.Unalias(v))
	case *types.Named:
		for a := range v.TypeArgs().Types() {
			walk(a)
		}

		p := v.Obj().Pkg()

		if p == nil {
			return
		}

		if p != self {
			into[p.Path()] = true

			return
		}

		walk(v.Underlying())

		for m := range v.Methods() {
			walk(m.Type())
		}
	case *types.Pointer:
		walk(v.Elem())
	case *types.Slice:
		walk(v.Elem())
	case *types.Array:
		walk(v.Elem())
	case *types.Chan:
		walk(v.Elem())
	case *types.Map:
		walk(v.Key())
		walk(v.Elem())
	case *types.Signature:
		walk(v.Params())
		walk(v.Results())

		for p := range v.TypeParams().TypeParams() {
			walk(p)
		}
	case *types.Tuple:
		for f := range v.Variables() {
			walk(f.Type())
		}
	case *types.Struct:
		for f := range v.Fields() {
			walk(f.Type())
		}
	case *types.Interface:
		for m := range v.ExplicitMethods() {
			walk(m.Type())
		}

		for e := range v.EmbeddedTypes() {
			walk(e)
		}
	case *types.Union:
		for term := range v.Terms() {
			walk(term.Type())
		}
	case *types.TypeParam:
		walk(v.Constraint())
	}
}
