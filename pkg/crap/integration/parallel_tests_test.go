package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

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
