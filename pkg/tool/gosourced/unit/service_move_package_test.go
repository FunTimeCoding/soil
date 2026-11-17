package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMovePackageBaseMismatch(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("package-move/src"))
	r, e := o.Packages.MovePackage(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"example/pkg/depot",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "name would change")
}

func TestMovePackageIntoItself(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("package-move/src"))
	r, e := o.Packages.MovePackage(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"example/pkg/outer/store/inner/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "into itself")
}
