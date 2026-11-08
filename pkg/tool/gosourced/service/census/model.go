package census

import "go/types"

func model(
	t types.Type,
	module string,
) *types.Named {
	if t == nil {
		return nil
	}

	n := named(t)

	if n == nil || !inModule(module, n.Obj().Pkg().Path()) {
		return nil
	}

	if _, okay := n.Underlying().(*types.Struct); !okay {
		return nil
	}

	return n
}
