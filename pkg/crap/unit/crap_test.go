package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/attribution"
	"github.com/funtimecoding/soil/pkg/crap/baseline"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/coverage"
	"github.com/funtimecoding/soil/pkg/crap/function"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/crap/mutation"
	"github.com/funtimecoding/soil/pkg/crap/score"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/notation"
	"path/filepath"
	"testing"
)

func TestMatrixSingleAndLoad(t *testing.T) {
	m := attribution.New()
	m.Add("p/TestA", []string{"m/a.go:1", "m/b.go:1"})
	m.Add("p/TestB", []string{"m/b.go:1", "m/c.go:1"})
	m.Add("p/TestC", []string{"m/b.go:1"})
	m.Add("p/TestD", nil)
	assert.Count(t, 4, m.Tests)
	assert.Count(t, 3, m.Covers)
	single := m.Single()
	assert.Count(t, 2, single)
	assert.String(t, "p/TestA", single["m/a.go:1"])
	assert.String(t, "p/TestB", single["m/c.go:1"])
	assert.Strings(t, []string{"m/a.go:1", "m/c.go:1"}, m.SingleKeys())
	load := m.Loads()
	assert.Count(t, 4, load)
	assert.String(t, "p/TestA", load[0].Test)
	assert.Integer(t, 1, load[0].Alone)
	assert.Integer(t, 2, load[0].Total)
	assert.String(t, "p/TestB", load[1].Test)
	assert.String(t, "p/TestC", load[2].Test)
	assert.Integer(t, 0, load[2].Alone)
	assert.String(t, "p/TestD", load[3].Test)
	assert.Integer(t, 0, load[3].Total)
}

func TestMatrixEmpty(t *testing.T) {
	m := attribution.New()
	assert.Count(t, 0, m.Single())
	assert.Count(t, 0, m.Loads())
}

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

func TestBuildPolicies(t *testing.T) {
	i := index.New("/r")
	i.Add(function.New("m", "/r/a.go", 1, "A", 4))
	i.Add(function.New("m", "/r/b.go", 1, "B", 4))
	coverageByKey := map[string]float64{"m/a.go:1": 50}
	pessimistic := crap.Build(i, coverageByKey, constant.Pessimistic)
	assert.Count(t, 2, pessimistic.Entries)
	assert.Float(t, 6, pessimistic.Entries[0].Score)
	assert.Float(t, 20, pessimistic.Entries[1].Score)
	assert.True(t, pessimistic.Entries[1].Missing)
	optimistic := crap.Build(i, coverageByKey, constant.Optimistic)
	assert.Float(t, 4, optimistic.Entries[1].Score)
	skip := crap.Build(i, coverageByKey, constant.Skip)
	assert.Count(t, 1, skip.Entries)
	assert.Integer(t, 1, skip.Skipped)
}

func TestReportOrderAndSummary(t *testing.T) {
	i := index.New("/r")
	i.Add(function.New("m", "/r/a.go", 1, "A", 2))
	i.Add(function.New("m", "/r/b.go", 1, "B", 6))
	r := crap.Build(
		i,
		map[string]float64{"m/a.go:1": 100, "m/b.go:1": 0},
		constant.Pessimistic,
	)
	sorted := r.Sorted()
	assert.String(t, "B", sorted[0].Function.Name)
	assert.Float(t, 42, sorted[0].Score)
	assert.Count(t, 1, r.Above(30))
	assert.Float(t, 44, r.Combined())
	assert.Float(t, 22, r.Average())
	assert.String(t, "b.go:1", sorted[0].Location("/r"))
}

func TestComplexityStraightLine(t *testing.T) {
	assert.Integer(t, 1, measure(t, `func f() { a := 1; _ = a }`))
}

func TestComplexityBranches(t *testing.T) {
	assert.Integer(
		t,
		4,
		measure(
			t,
			`func f(a, b bool) {
	if a { }
	for i := 0; i < 2; i++ { }
	for range []int{} { }
}`,
		),
	)
}

func TestComplexitySwitchAndSelect(t *testing.T) {
	assert.Integer(
		t,
		4,
		measure(
			t,
			`func f(a int, c chan int) {
	switch a {
	case 1:
	case 2:
	default:
	}
	select {
	case <-c:
	default:
	}
}`,
		),
	)
}

func TestComplexityBooleanOperators(t *testing.T) {
	assert.Integer(
		t,
		3,
		measure(t, `func f(a, b, c bool) bool { return a && b || c }`),
	)
}

func TestParseFunctionReport(t *testing.T) {
	v := coverage.Parse(
		`github.com/x/m/pkg/a/add.go:3:		Add		100.0%
github.com/x/m/pkg/a/run.go:10:		Run		0.0%
github.com/x/m/pkg/a/walk.go:11:		walk		95.2%
total:			(statements)		87.9%
`,
	)
	assert.Count(t, 3, v)
	assert.Float(t, 100, v["github.com/x/m/pkg/a/add.go:3"])
	assert.Float(t, 0, v["github.com/x/m/pkg/a/run.go:10"])
	assert.Float(t, 95.2, v["github.com/x/m/pkg/a/walk.go:11"])
}

func TestParseEmpty(t *testing.T) {
	assert.Count(t, 0, coverage.Parse(""))
}

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

func TestScoreUncovered(t *testing.T) {
	assert.Float(t, 30, score.Score(5, 0))
	assert.Float(t, 72, score.Score(8, 0))
}

func TestScoreCovered(t *testing.T) {
	assert.Float(t, 5, score.Score(5, 100))
}

func TestScorePartial(t *testing.T) {
	assert.Float(t, 5.390625, score.Score(5, 75))
}

func TestParsePolicy(t *testing.T) {
	assert.True(t, score.ParsePolicy("") == constant.Pessimistic)
	assert.True(t, score.ParsePolicy("Optimistic") == constant.Optimistic)
	assert.True(t, score.ParsePolicy("SKIP") == constant.Skip)
}

func TestParsePolicyUnknownPanics(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	score.ParsePolicy("hopeful")
}
