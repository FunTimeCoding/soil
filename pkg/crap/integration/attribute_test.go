package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/coverage"
	"github.com/funtimecoding/soil/pkg/crap/function"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
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
