package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMatchPatternSpreadHoleType(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("census/src"),
	)
	s := service_tester.NewService(t)
	r, match, e := s.MatchPattern(
		d,
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
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("census/src"),
	)
	s := service_tester.NewService(t)
	r, inventory, e := s.ListCalls(d, "example/pkg/missing", 0)
	assert.FatalOnError(t, e)
	assert.True(t, inventory == nil)
	testutil.AssertBlockedContains(t, r, constant.MissingPackage)
}
