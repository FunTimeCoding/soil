package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/function"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/crap/mutation"
	"testing"
)

func TestMutationLoadAndByFile(t *testing.T) {
	m := loadGremlins(t)
	assert.String(t, "example.test/m", m.Module)
	byFile := m.ByFile()
	assert.Count(t, 2, byFile)
	assert.Count(t, 2, byFile["example.test/m/pkg/a/half.go"])
	assert.Count(
		t,
		1,
		mutation.Within(byFile["example.test/m/pkg/a/half.go"], 5, 9),
	)
}

func TestAnnotateRescoresSurvivors(t *testing.T) {
	i := index.New("/r")
	half := function.New(
		"example.test/m/pkg/a",
		"/r/pkg/a/half.go",
		3,
		"Half",
		2,
	)
	half.EndLine = 9
	covered := function.New(
		"example.test/m/pkg/a",
		"/r/pkg/a/covered.go",
		3,
		"Covered",
		2,
	)
	covered.EndLine = 9
	untested := function.New(
		"example.test/m/pkg/a",
		"/r/pkg/a/thing.go",
		5,
		"Untested",
		3,
	)
	untested.EndLine = 11
	i.Add(half)
	i.Add(covered)
	i.Add(untested)
	r := crap.Build(
		i,
		map[string]float64{
			half.Key():     100,
			covered.Key():  100,
			untested.Key(): 0,
		},
		constant.Pessimistic,
	)
	crap.Annotate(r, loadGremlins(t))
	byName := map[string]float64{}
	untrusted := map[string]bool{}

	for _, e := range r.Entries {
		byName[e.Function.Name] = e.Score
		untrusted[e.Function.Name] = e.Untrusted
	}

	assert.Float(t, 6, byName["Half"])
	assert.True(t, untrusted["Half"])
	assert.Float(t, 2, byName["Covered"])
	assert.False(t, untrusted["Covered"])
	assert.Float(t, 12, byName["Untested"])
	assert.False(t, untrusted["Untested"])
}

func TestKillRatio(t *testing.T) {
	f := function.New("m", "/r/a.go", 1, "A", 1)
	f.EndLine = 10
	i := index.New("/r")
	i.Add(f)
	r := crap.Build(i, map[string]float64{f.Key(): 100}, constant.Pessimistic)
	e := r.Entries[0]
	assert.Float(t, 1, e.KillRatio())
	e.Annotate(
		[]*mutation.Mutant{
			mutation.NewMutant(constant.MutantKilled, 2),
			mutation.NewMutant(constant.MutantKilled, 3),
			mutation.NewMutant(constant.MutantLived, 4),
			mutation.NewMutant("NOT COVERED", 5),
		},
	)
	assert.Integer(t, 2, e.Killed)
	assert.Count(t, 1, e.Lived)
	assert.True(t, e.KillRatio() > 0.66 && e.KillRatio() < 0.67)
}
