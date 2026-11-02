package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMovePackageBaseMismatch(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("package-move/src"),
	)
	s := service_tester.NewService(t)
	r, e := s.MovePackage(
		t.Context(),
		d,
		"example/pkg/outer/store",
		"example/pkg/depot",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "name would change")
}

func TestMovePackageIntoItself(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("package-move/src"),
	)
	s := service_tester.NewService(t)
	r, e := s.MovePackage(
		t.Context(),
		d,
		"example/pkg/outer/store",
		"example/pkg/outer/store/inner/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "into itself")
}
