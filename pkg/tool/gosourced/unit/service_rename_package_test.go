package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenamePackageInvalidName(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("package-rename/src"))
	r, e := o.Packages.RenamePackage(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"9depot",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not a valid")
}
