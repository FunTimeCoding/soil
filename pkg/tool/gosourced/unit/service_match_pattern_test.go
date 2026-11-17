package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMatchPatternSpreadHoleType(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("pattern/src"))
	r, match, e := o.Pattern.MatchPattern(
		o.Directory,
		"example/pkg/client",
		"Client",
		"",
		"func pattern(c *client.Client, x []int) {\n\tif c.Client(x...) {\n\t}\n}",
	)
	assert.FatalOnError(t, e)
	assert.True(t, match == nil)
	testutil.AssertBlockedContains(t, r, "[]any")
}

func TestListCallsUnknownRegion(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("pattern/src"))
	r, inventory, e := o.Reference.ListCalls(
		o.Directory,
		"example/pkg/missing",
		0,
	)
	assert.FatalOnError(t, e)
	assert.True(t, inventory == nil)
	testutil.AssertBlockedContains(t, r, constant.MissingPackage)
}
