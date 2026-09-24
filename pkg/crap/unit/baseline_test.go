package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/baseline"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/notation"
	"path/filepath"
	"testing"
)

func TestBaselineRoundTrip(t *testing.T) {
	root := t.TempDir()
	previous := build(map[string]float64{"m/a.go:1": 100, "m/b.go:1": 0})
	testutil.WriteFile(
		t,
		root,
		"baseline.json",
		notation.MarshalIndent(previous),
	)
	b := baseline.Load(filepath.Join(root, "baseline.json"))
	assert.Float(t, 44, b.Combined)
	v, okay := b.Score("m/b.go:1")
	assert.True(t, okay)
	assert.Float(t, 42, v)
	_, okay = b.Score("m/c.go:1")
	assert.False(t, okay)
}

func TestCompareMarksDeltas(t *testing.T) {
	previous := baseline.New(
		build(map[string]float64{"m/a.go:1": 100, "m/b.go:1": 0}),
	)
	current := build(map[string]float64{"m/a.go:1": 50, "m/b.go:1": 100})
	current.Add(entryFor("m", "/r/c.go", "C", 3, 0))
	crap.Compare(current, previous)
	byName := map[string]float64{}
	isNew := map[string]bool{}

	for _, e := range current.Entries {
		byName[e.Function.Name] = e.Delta()
		isNew[e.Function.Name] = e.IsNew()
	}

	assert.Float(t, 0.5, byName["A"])
	assert.Float(t, -36, byName["B"])
	assert.True(t, isNew["C"])
	assert.Float(t, 12, byName["C"])
	regressed := current.Regressions(0.01, false)
	assert.Count(t, 1, regressed)
	assert.String(t, "A", regressed[0].Function.Name)
}

func TestRegressionsIgnoreCovered(t *testing.T) {
	previous := baseline.New(build(map[string]float64{"m/a.go:1": 100}))
	current := build(map[string]float64{"m/a.go:1": 100})
	current.Entries[0].Function.Complexity = 5
	current.Entries[0].Score = 5
	crap.Compare(current, previous)
	assert.Count(t, 1, current.Regressions(0.01, false))
	assert.Count(t, 0, current.Regressions(0.01, true))
	assert.Count(t, 0, current.Regressions(10, false))
}

func TestUnchangedWithinTolerance(t *testing.T) {
	previous := baseline.New(build(map[string]float64{"m/a.go:1": 100}))
	current := build(map[string]float64{"m/a.go:1": 100})
	crap.Compare(current, previous)
	assert.True(t, current.Entries[0].Unchanged(0.01))
	assert.False(t, current.Entries[0].Regressed(0.01))
}
