package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/coverage"
	"github.com/funtimecoding/soil/pkg/crap/function"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/crap/option"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestAttributeEndToEnd(t *testing.T) {
	root := module(t)
	testutil.WriteFile(
		t,
		root,
		"pkg/a/unit/b_test.go",
		`package unit

import (
	"example.test/m/pkg/a"
	"testing"
)

func TestHalfOnly(t *testing.T) {
	a.Half(-1)
}

func TestNothing(t *testing.T) {}
`,
	)
	packages := coverage.TestPackages(root, constant.AllPackages)
	assert.Strings(t, []string{"example.test/m/pkg/a/unit"}, packages)
	i := index.Load(root, constant.AllPackages)
	m := crap.Attribute(root, i, packages, []string{constant.AllPackages})
	assert.Strings(
		t,
		[]string{
			"example.test/m/pkg/a/unit/TestCovered",
			"example.test/m/pkg/a/unit/TestHalfOnly",
			"example.test/m/pkg/a/unit/TestNothing",
		},
		m.Tests,
	)
	covered := function.NewKey("example.test/m/pkg/a", "covered.go", 3)
	half := function.NewKey("example.test/m/pkg/a", "half.go", 3)
	assert.Strings(
		t,
		[]string{"example.test/m/pkg/a/unit/TestCovered"},
		m.Covers[covered],
	)
	assert.Count(t, 2, m.Covers[half])
	single := m.Single()
	assert.Count(t, 1, single)
	assert.String(t, "example.test/m/pkg/a/unit/TestCovered", single[covered])
	load := m.Loads()
	assert.String(t, "example.test/m/pkg/a/unit/TestCovered", load[0].Test)
	assert.Integer(t, 1, load[0].Alone)
	assert.Integer(t, 2, load[0].Total)
	assert.Integer(t, 0, load[2].Total)
}

func TestCompileAndListTests(t *testing.T) {
	root := module(t)
	binary := coverage.Compile(
		root,
		t.TempDir(),
		"example.test/m/pkg/a/unit",
		[]string{constant.AllPackages},
	)
	assert.True(t, system.IsExecutable(binary))
	assert.Strings(t, []string{"TestCovered"}, coverage.ListTests(binary))
	profile := coverage.RunTestAlone(root, binary, "TestCovered", t.TempDir())
	keys := coverage.CoveredKeys(coverage.Functions(root, profile))
	assert.Count(t, 2, keys)
}

func TestCoverageEndToEnd(t *testing.T) {
	root := module(t)
	profile := coverage.Run(root, "", constant.AllPackages)
	assert.True(t, system.FileExists(profile))
	r := crap.Build(
		index.Load(root, constant.AllPackages),
		coverage.Functions(root, profile),
		constant.Pessimistic,
	)
	assert.Count(t, 3, r.Entries)
	byName := map[string]float64{}
	missing := map[string]bool{}

	for _, e := range r.Entries {
		byName[e.Function.Name] = e.Coverage
		missing[e.Function.Name] = e.Missing
	}

	assert.Float(t, 100, byName["Covered"])
	assert.True(t, byName["Half"] > 0 && byName["Half"] < 100)
	assert.Float(t, 0, byName["Untested"])
	assert.False(t, missing["Untested"])
	assert.Float(t, 12, r.Sorted()[0].Score)
}

func TestIndexLoad(t *testing.T) {
	root := module(t)
	i := index.Load(root, constant.AllPackages)
	assert.Count(t, 3, i.Functions)
	half := i.ByKey(function.NewKey("example.test/m/pkg/a", "half.go", 3))
	assert.NotNil(t, half)
	assert.String(t, "Half", half.Name)
	assert.Integer(t, 2, half.Complexity)
	assert.String(t, "example.test/m/pkg/a", half.Package)
	untested := i.ByKey(function.NewKey("example.test/m/pkg/a", "thing.go", 5))
	assert.String(t, "*Thing", untested.Receiver)
	assert.String(t, "*Thing.Untested", untested.QualifiedName())
	assert.Integer(t, 3, untested.Complexity)
	assert.Integer(t, 11, untested.EndLine)
}

func TestKeyMatchesCoverReport(t *testing.T) {
	assert.String(
		t,
		"example.test/m/pkg/a/half.go:3",
		function.NewKey("example.test/m/pkg/a", "/any/where/pkg/a/half.go", 3),
	)
}

func TestParallelTestsAbsent(t *testing.T) {
	assert.Count(t, 0, index.ParallelTests(module(t), constant.AllPackages))
}

func TestParallelTestsDetected(t *testing.T) {
	root := module(t)
	testutil.WriteFile(
		t,
		root,
		"pkg/a/unit/parallel_test.go",
		`package unit

import "testing"

func TestParallel(t *testing.T) {
	t.Parallel()
}
`,
	)
	assert.Strings(
		t,
		[]string{"example.test/m/pkg/a/unit"},
		index.ParallelTests(root, constant.AllPackages),
	)
}

func TestRunAttributeRefusesUnscoped(t *testing.T) {
	o := option.NewAttribute()
	o.Root = module(t)
	o.Patterns = []string{constant.AllPackages}
	out := assert.Capture(
		t,
		func() { assert.Integer(t, 1, crap.RunAttribute(o)) },
	)
	assert.StringContains(t, "name the packages", out)
}

func TestRunAttributeScopedPasses(t *testing.T) {
	o := option.NewAttribute()
	o.Root = module(t)
	o.Patterns = []string{"./pkg/a/..."}
	o.Notation = true
	assert.Capture(t, func() { assert.Integer(t, 0, crap.RunAttribute(o)) })
}

func TestRunFailAboveThreshold(t *testing.T) {
	root := module(t)
	o := report(root)
	o.FailAbove = true
	o.Threshold = 100
	assert.Integer(t, 0, crap.Run(o))
	o.Threshold = 10
	assert.Integer(t, 1, crap.Run(o))
}

func TestRunFailRegression(t *testing.T) {
	root := module(t)
	o := report(root)
	o.Notation = true
	o.Profile = ""
	r := assert.Capture(t, func() { crap.Run(o) })
	testutil.WriteFile(t, root, "baseline.json", r)
	o.Notation = false
	o.Baseline = filepath.Join(root, "baseline.json")
	o.FailRegression = true
	assert.Integer(t, 0, crap.Run(o))
	testutil.WriteFile(
		t,
		root,
		"pkg/a/covered.go",
		`package a

func Covered(a bool) int {
	if a {
		return 1
	}

	if !a {
		return 3
	}

	return 2
}
`,
	)
	assert.Integer(t, 1, crap.Run(o))
	o.IgnoreCovered = true
	assert.Integer(t, 1, crap.Run(o))
}

func TestRunAttributeRefusesParallel(t *testing.T) {
	root := module(t)
	testutil.WriteFile(
		t,
		root,
		"pkg/a/unit/parallel_test.go",
		"package unit\n\nimport \"testing\"\n\nfunc TestParallel(t *testing.T) { t.Parallel() }\n",
	)
	o := option.NewAttribute()
	o.Root = root
	o.Patterns = []string{"./pkg/a/..."}
	out := assert.Capture(
		t,
		func() { assert.Integer(t, 1, crap.RunAttribute(o)) },
	)
	assert.StringContains(t, "parallel subtests", out)
}

func TestRunAttributeNotation(t *testing.T) {
	root := module(t)
	o := option.NewAttribute()
	o.Root = root
	o.Patterns = []string{"./pkg/a/..."}
	o.Notation = true
	out := assert.Capture(
		t,
		func() { assert.Integer(t, 0, crap.RunAttribute(o)) },
	)
	var decoded map[string]any
	notation.MustDecode(out, &decoded, false)
	assert.MapHasKey(t, decoded, "tests")
	assert.MapHasKey(t, decoded, "covers")
}
