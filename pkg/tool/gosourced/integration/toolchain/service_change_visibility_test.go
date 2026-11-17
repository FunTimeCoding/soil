package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestUnexportFunction(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("unexport-function/src"),
	)
	r, e := o.Symbol.ChangeVisibility(
		t.Context(),
		o.Directory,
		"IsGenerated",
		"example/pkg/target",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	helper := o.Read("pkg/target/is_generated.go")
	assert.StringContains(t, "func isGenerated(", helper)
	run := o.Read("pkg/target/run.go")
	assert.StringContains(t, "isGenerated(name)", run)
}

func TestExportFunction(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("export-function/src"),
	)
	r, e := o.Symbol.ChangeVisibility(
		t.Context(),
		o.Directory,
		"isGenerated",
		"example/pkg/target",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	helper := o.Read("pkg/target/is_generated.go")
	assert.StringContains(t, "func IsGenerated(", helper)
	run := o.Read("pkg/target/run.go")
	assert.StringContains(t, "IsGenerated(\"test\")", run)
}

func TestUnexportMethod(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("unexport-method/src"),
	)
	r, e := o.Symbol.ChangeVisibility(
		t.Context(),
		o.Directory,
		"Save",
		"example/pkg/target",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	save := o.Read("pkg/target/save.go")
	assert.StringContains(t, "func (s *Store) save(", save)
	run := o.Read("pkg/target/run.go")
	assert.StringContains(t, "v.save(\"test\")", run)
}

func TestCollisionDetection(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("collision/src"))
	r, e := o.Symbol.ChangeVisibility(
		t.Context(),
		o.Directory,
		"IsGenerated",
		"example/pkg/target",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestSymbolNotFound(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("unexport-function/src"),
	)
	r, e := o.Symbol.ChangeVisibility(
		t.Context(),
		o.Directory,
		"Missing",
		"example/pkg/target",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestPackageNotFound(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("unexport-function/src"),
	)
	r, e := o.Symbol.ChangeVisibility(
		t.Context(),
		o.Directory,
		"IsGenerated",
		"example/pkg/missing",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestUnexportBlockedByCrossPackageCaller(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("cross-package/src"))
	r, e := o.Symbol.ChangeVisibility(
		t.Context(),
		o.Directory,
		"IsValid",
		"example/pkg/target",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "would lose access")
	helper := o.Read("pkg/target/is_valid.go")
	assert.StringContains(t, "func IsValid(", helper)
}
