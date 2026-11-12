package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/kind_tester"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestIndexStoresEverythingUnderOneGeneration(t *testing.T) {
	directory := kind_tester.Module(t)
	store := t.TempDir()
	var extracted, external int
	alfa := kind_tester.Counting("alfa", &extracted, &external)
	index.Facts[*string](index.New(store, directory, alfa), alfa)
	entries, e := os.ReadDir(store)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, entries)
	assert.True(t, strings.HasPrefix(entries[0].Name(), runtime.Version()))
}

func TestIndexKeepsEachBuildApartInsideItsGeneration(t *testing.T) {
	directory := kind_tester.Module(t)
	store := t.TempDir()
	var extracted, external int
	alfa := kind_tester.Counting("alfa", &extracted, &external)
	index.Facts[*string](
		index.New(store, directory, alfa).WithFingerprint("first"),
		alfa,
	)
	index.Facts[*string](
		index.New(store, directory, alfa).WithFingerprint("second"),
		alfa,
	)
	generations, e := os.ReadDir(store)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, generations)
	root := filepath.Join(
		store,
		generations[0].Name(),
		constant.IndexBuildDirectory,
	)
	builds, f := os.ReadDir(root)
	assert.FatalOnError(t, f)
	holding := 0

	for _, b := range builds {
		if _, g := os.Stat(
			filepath.Join(root, b.Name(), constant.IndexFactsKind),
		); g == nil {
			holding++
		}
	}

	assert.Integer(t, 2, holding)
}

func TestIndexLeavesForeignFileInItsRoot(t *testing.T) {
	directory := kind_tester.Module(t)
	store := t.TempDir()
	var extracted, external int
	alfa := kind_tester.Counting("alfa", &extracted, &external)
	foreign := filepath.Join(store, "gosourced.yaml")
	assert.FatalOnError(t, os.WriteFile(foreign, nil, 0644))
	old := time.Now().Add(-30 * 24 * time.Hour)
	assert.FatalOnError(t, os.Chtimes(foreign, old, old))
	index.Facts[*string](index.New(store, directory, alfa), alfa)
	index.Prune(store)
	_, e := os.Stat(foreign)
	assert.FatalOnError(t, e)
}

func TestIndexRemovesFlatLayoutPastRetention(t *testing.T) {
	directory := kind_tester.Module(t)
	store := t.TempDir()
	var extracted, external int
	alfa := kind_tester.Counting("alfa", &extracted, &external)
	flat := filepath.Join(store, constant.IndexPackageKind)
	assert.FatalOnError(t, os.MkdirAll(flat, 0755))
	old := time.Now().Add(-4 * 24 * time.Hour)
	assert.FatalOnError(t, os.Chtimes(flat, old, old))
	index.Facts[*string](index.New(store, directory, alfa), alfa)
	_, e := os.Stat(flat)
	assert.True(t, os.IsNotExist(e))
}

func TestIndexPruneLeavesOnlyCurrentGeneration(t *testing.T) {
	directory := kind_tester.Module(t)
	store := t.TempDir()
	var extracted, external int
	alfa := kind_tester.Counting("alfa", &extracted, &external)
	assert.FatalOnError(
		t,
		os.MkdirAll(filepath.Join(store, "go1.0-0123456789ab"), 0755),
	)
	index.Facts[*string](index.New(store, directory, alfa), alfa)
	assert.Any(t, []string{"go1.0-0123456789ab"}, index.Prune(store))
	entries, e := os.ReadDir(store)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, entries)
	assert.True(t, strings.HasPrefix(entries[0].Name(), runtime.Version()))
}
