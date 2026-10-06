package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/index"
	"maps"
	"slices"
	"testing"
)

func TestIndexComputesAKindOnlyWhenRequested(t *testing.T) {
	directory := kindModule(t)
	store := t.TempDir()
	var alfaExtracted, alfaExternal, bravoExtracted, bravoExternal int
	alfa := countingKind("alfa", &alfaExtracted, &alfaExternal)
	bravo := countingKind("bravo", &bravoExtracted, &bravoExternal)
	first := index.Facts[*string](index.New(store, directory, alfa), alfa)
	assert.Strings(
		t,
		[]string{"example/one", "example/two"},
		slices.Sorted(maps.Keys(first)),
	)
	assert.String(t, "example/one", *first["example/one"])
	assert.Integer(t, 2, alfaExtracted)
	assert.Integer(t, 0, bravoExtracted)
	index.Facts[*string](index.New(store, directory, alfa), alfa)
	assert.Integer(t, 2, alfaExtracted)
	both := index.New(store, directory, alfa, bravo)
	assert.Integer(t, 2, len(index.Facts[*string](both, bravo)))
	assert.Integer(t, 2, alfaExtracted)
	assert.Integer(t, 2, bravoExtracted)
}

func TestIndexKeepsExternalsPerKind(t *testing.T) {
	directory := kindModule(t)
	store := t.TempDir()
	var extracted, external int
	alfa := countingKind("alfa", &extracted, &external)
	w := index.New(store, directory, alfa)
	var paths []string

	for _, p := range index.Externals[*string](w, alfa) {
		paths = append(paths, *p)
	}

	assert.True(t, slices.Contains(paths, "strings"))
	stored := external
	index.Externals[*string](index.New(store, directory, alfa), alfa)
	assert.Integer(t, stored, external)
}

func TestIndexReferencesNeedNoKind(t *testing.T) {
	directory := kindModule(t)
	r := index.New(t.TempDir(), directory).References()
	assert.True(t, r.Has("example/two"))
}
