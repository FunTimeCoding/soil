package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestApplyPatternUnknownHole(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("pattern/src"))
	r, apply, e := o.Pattern.ApplyPattern(
		t.Context(),
		o.Directory,
		"example/pkg/pair",
		"Compare",
		"",
		"func pattern(x int, y int) int {\n\treturn pair.Compare(x, y)\n}",
		"func replacement(x int, z int) int {\n\treturn pair.Compare(z, x)\n}",
		true,
		true,
	)
	assert.FatalOnError(t, e)
	assert.True(t, apply == nil)
	testutil.AssertBlockedContains(t, r, "z")
}
