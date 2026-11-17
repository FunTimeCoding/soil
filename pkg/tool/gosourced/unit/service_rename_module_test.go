package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameModuleSamePath(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("module-rename/src"))
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/lib/v2",
		"example/lib/v2",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "same")
}
