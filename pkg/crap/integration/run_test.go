package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/option"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/notation"
	"path/filepath"
	"testing"
)

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
