package unit

import "go/types"

func aliasedNode() *types.Alias {
	inner := types.NewNamed(
		types.NewTypeName(
			0,
			types.NewPackage("example.org/alfa/internal/codec", "codec"),
			"Node",
			nil,
		),
		types.NewStruct(nil, nil),
		nil,
	)

	return types.NewAlias(
		types.NewTypeName(
			0,
			types.NewPackage("example.org/alfa/codec", "codec"),
			"Node",
			nil,
		),
		inner,
	)
}
