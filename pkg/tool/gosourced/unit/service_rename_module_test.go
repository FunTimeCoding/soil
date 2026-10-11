package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameModuleSamePath(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("module-rename/src"),
	)
	s := service_tester.NewService()
	r, e := s.RenameModule(d, "example/lib/v2", "example/lib/v2", false, false)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "same")
}
