package unit

import "go/types"

func nestedScope() *types.Scope {
	p := types.NewScope(types.Universe, 0, 1, "package scope")
	function := types.NewScope(p, 0, 1, "function")
	block := types.NewScope(function, 0, 1, "block")
	p.Insert(types.NewVar(0, nil, "p", types.Typ[types.Int]))
	function.Insert(types.NewVar(0, nil, "f", types.Typ[types.Int]))
	block.Insert(types.NewVar(0, nil, "b", types.Typ[types.Int]))

	return types.NewScope(block, 0, 1, "inner")
}
