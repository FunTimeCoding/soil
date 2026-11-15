package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/canonical_type"
	"github.com/funtimecoding/soil/pkg/source/unit/canonical_tester"
	"go/types"
	"testing"
)

func TestCanonicalTypeResolvesAliasBehindPointer(t *testing.T) {
	assert.String(
		t,
		"*example.org/alfa/internal/codec.Node",
		canonical_type.String(
			types.NewPointer(canonical_tester.AliasedNode()),
			(*types.Package).Path,
		),
	)
}

func TestCanonicalTypeResolvesAliasInsideCollections(t *testing.T) {
	assert.String(
		t,
		"map[string][]example.org/alfa/internal/codec.Node",
		canonical_type.String(
			types.NewMap(
				types.Typ[types.String],
				types.NewSlice(canonical_tester.AliasedNode()),
			),
			(*types.Package).Path,
		),
	)
}

func TestCanonicalTypeSpellsAnyAsEmptyInterface(t *testing.T) {
	assert.String(
		t,
		canonical_type.String(
			types.NewInterfaceType(nil, nil),
			(*types.Package).Path,
		),
		canonical_type.String(
			types.Universe.Lookup("any").Type(),
			(*types.Package).Path,
		),
	)
}

func TestCanonicalTypeResolvesInlineStructField(t *testing.T) {
	assert.String(
		t,
		"struct{N example.org/alfa/internal/codec.Node \"yaml:\\\"n\\\"\"}",
		canonical_type.String(
			types.NewStruct(
				[]*types.Var{
					types.NewField(
						0,
						nil,
						"N",
						canonical_tester.AliasedNode(),
						false,
					),
				},
				[]string{"yaml:\"n\""},
			),
			(*types.Package).Path,
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
									types.NewPointer(
										canonical_tester.AliasedNode(),
									),
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
			(*types.Package).Path,
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
					types.NewParam(
						0,
						nil,
						"",
						types.NewPointer(canonical_tester.AliasedNode()),
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
			(*types.Package).Path,
		),
	)
}
