package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/coverage"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

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
