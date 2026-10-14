package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/kind_tester"
	"maps"
	"slices"
	"strings"
	"testing"
)

func TestIndexComputesAKindOnlyWhenRequested(t *testing.T) {
	directory := kind_tester.Module(t)
	store := t.TempDir()
	var alfaExtracted, alfaExternal, bravoExtracted, bravoExternal int
	alfa := kind_tester.Counting("alfa", &alfaExtracted, &alfaExternal)
	bravo := kind_tester.Counting("bravo", &bravoExtracted, &bravoExternal)
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
	directory := kind_tester.Module(t)
	store := t.TempDir()
	var extracted, external int
	alfa := kind_tester.Counting("alfa", &extracted, &external)
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
	directory := kind_tester.Module(t)
	r := index.New(t.TempDir(), directory).References()
	assert.True(t, r.Has("example/two"))
}

func TestIndexExtractsSyntaxExternalsTransitively(t *testing.T) {
	directory := kind_tester.Module(t)
	store := t.TempDir()
	var external int
	alfa := kind_tester.CountingSyntax("alfa", &external)
	var paths []string
	var bare []string

	for _, p := range index.Externals[*string](
		index.New(store, directory, alfa),
		alfa,
	) {
		if strings.HasPrefix(*p, "!") {
			bare = append(bare, *p)
		}

		paths = append(paths, *p)
	}

	assert.True(t, slices.Contains(paths, "strings"))
	assert.True(t, slices.Contains(paths, "unicode/utf8"))
	assert.Strings(t, []string{"!unsafe"}, bare)
	stored := external
	index.Externals[*string](index.New(store, directory, alfa), alfa)
	assert.Integer(t, stored, external)
}

func TestIndexRecomputesFactsForAnotherBuild(t *testing.T) {
	directory := kind_tester.Module(t)
	store := t.TempDir()
	var extracted, external int
	alfa := kind_tester.Counting("alfa", &extracted, &external)
	index.Facts[*string](
		index.New(store, directory, alfa).WithFingerprint("first"),
		alfa,
	)
	assert.Integer(t, 2, extracted)
	index.Facts[*string](
		index.New(store, directory, alfa).WithFingerprint("first"),
		alfa,
	)
	assert.Integer(t, 2, extracted)
	index.Facts[*string](
		index.New(store, directory, alfa).WithFingerprint("second"),
		alfa,
	)
	assert.Integer(t, 4, extracted)
}
