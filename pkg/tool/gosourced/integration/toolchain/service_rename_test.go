package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameFunction(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 3)
	declaration := o.Read("pkg/target/is_generated.go")
	assert.StringContains(t, "func IsGenerated(", declaration)
	samePackage := o.Read("pkg/target/check.go")
	assert.StringContains(t, "IsGenerated(content)", samePackage)
	crossPackage := o.Read("pkg/caller/run.go")
	assert.StringContains(t, "target.IsGenerated(content)", crossPackage)
}

func TestRenameFunctionSamePackage(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("unexport-function/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGenerated",
		"WasGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	declaration := o.Read("pkg/target/was_generated.go")
	assert.StringContains(t, "func WasGenerated(", declaration)
	run := o.Read("pkg/target/run.go")
	assert.StringContains(t, "WasGenerated(name)", run)
}

func TestRenameMethod(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-method/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"FindByName",
		"LookupByName",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 3)
	method := o.Read("pkg/target/lookup_by_name.go")
	assert.StringContains(t, "func (s *Store) LookupByName(", method)
	samePackage := o.Read("pkg/target/run.go")
	assert.StringContains(t, "v.LookupByName(\"test\")", samePackage)
	crossPackage := o.Read("pkg/caller/run.go")
	assert.StringContains(t, "s.LookupByName(\"test\")", crossPackage)
}

func TestRenameToUnexportedBlockedByCrossPackage(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-unexport/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Validate",
		"check",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "would lose access")
	source := o.Read("pkg/target/validate.go")
	assert.StringContains(t, "func Validate(", source)
}

func TestRenameToUnexportedAllowedWithinPackage(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("unexport-function/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGenerated",
		"wasGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	declaration := o.Read("pkg/target/was_generated.go")
	assert.StringContains(t, "func wasGenerated(", declaration)
}

func TestRenameCollision(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"Check",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestRenameSymbolNotFound(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Missing",
		"Something",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestRenamePackageNotFound(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/missing",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestRenameReceiverNotFound(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-method/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"FindByName",
		"LookupByName",
		"Missing",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestRenameMethodNotFound(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-method/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Missing",
		"Something",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestRenameMethodCollision(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("method-collision/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Save",
		"Load",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestRenameSameName(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGeneratedHeader",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestRenameLongerKeepsComments(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-longer/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target/constant",
		"Legend",
		"TargetLegendExtended",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.Formatted(t, source)
	assert.StringContains(t, "TargetLegendExtended = \"--legend\"", source)
	assert.StringContains(t, "// alfa trailing", source)
	assert.StringContains(t, "// bold trailing", source)
	assert.StringContains(t, "// wide trailing", source)
	run := o.Read("pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "TargetLegendExtended", run)
}
