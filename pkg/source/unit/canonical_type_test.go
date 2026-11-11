package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/canonical_type"
	"go/types"
	"testing"
)

func TestCanonicalTypeResolvesAliasBehindPointer(t *testing.T) {
	assert.String(
		t,
		"*example.org/alfa/internal/codec.Node",
		canonical_type.String(types.NewPointer(aliasedNode()), qualifiedPath),
	)
}

func TestCanonicalTypeResolvesAliasInsideCollections(t *testing.T) {
	assert.String(
		t,
		"map[string][]example.org/alfa/internal/codec.Node",
		canonical_type.String(
			types.NewMap(
				types.Typ[types.String],
				types.NewSlice(aliasedNode()),
			),
			qualifiedPath,
		),
	)
}

func TestCanonicalTypeSpellsAnyAsEmptyInterface(t *testing.T) {
	assert.String(
		t,
		canonical_type.String(types.NewInterfaceType(nil, nil), qualifiedPath),
		canonical_type.String(
			types.Universe.Lookup("any").Type(),
			qualifiedPath,
		),
	)
}

func TestCanonicalTypeResolvesInlineStructField(t *testing.T) {
	assert.String(
		t,
		"struct{N example.org/alfa/internal/codec.Node \"yaml:\\\"n\\\"\"}",
		canonical_type.String(
			types.NewStruct(
				[]*types.Var{types.NewField(0, nil, "N", aliasedNode(), false)},
				[]string{"yaml:\"n\""},
			),
			qualifiedPath,
		),
	)
}

func TestCanonicalTypeResolvesInlineInterfaceMethod(t *testing.T) {
	assert.String(
		t,
		"interface{Decode(*example.org/alfa/internal/codec.Node) error}",
		canonical_type.String(
			types.NewInterfaceType(
				[]*types.Func{
					types.NewFunc(
						0,
						nil,
						"Decode",
						types.NewSignatureType(
							nil,
							nil,
							nil,
							types.NewTuple(
								types.NewParam(
									0,
									nil,
									"",
									types.NewPointer(aliasedNode()),
								),
							),
							types.NewTuple(
								types.NewParam(
									0,
									nil,
									"",
									types.Universe.Lookup("error").Type(),
								),
							),
							false,
						),
					),
				},
				nil,
			).Complete(),
			qualifiedPath,
		),
	)
}

func TestCanonicalTypeResolvesSignatureParameters(t *testing.T) {
	assert.String(
		t,
		"func(*example.org/alfa/internal/codec.Node) error",
		canonical_type.String(
			types.NewSignatureType(
				nil,
				nil,
				nil,
				types.NewTuple(
					types.NewParam(0, nil, "", types.NewPointer(aliasedNode())),
				),
				types.NewTuple(
					types.NewParam(
						0,
						nil,
						"",
						types.Universe.Lookup("error").Type(),
					),
				),
				false,
			),
			qualifiedPath,
		),
	)
}
