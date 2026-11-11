package canonical_type

import "go/types"

func unaliasedInterface(i *types.Interface) types.Type {
	if i.IsImplicit() {
		return i
	}

	methods := make([]*types.Func, i.NumExplicitMethods())

	for j := range i.NumExplicitMethods() {
		m := i.ExplicitMethod(j)
		s, okay := unaliased(m.Type()).(*types.Signature)

		if !okay {
			return i
		}

		methods[j] = types.NewFunc(m.Pos(), m.Pkg(), m.Name(), s)
	}

	embedded := make([]types.Type, i.NumEmbeddeds())

	for j := range i.NumEmbeddeds() {
		embedded[j] = unaliased(i.EmbeddedType(j))
	}

	return types.NewInterfaceType(methods, embedded).Complete()
}
