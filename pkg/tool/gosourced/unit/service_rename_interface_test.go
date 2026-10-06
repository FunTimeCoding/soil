package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameRefusesAMethodSatisfyingAnInterface(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface/src"),
	)
	before := treeOf(t, d)
	r, e := testService().Rename(
		d,
		"example/pkg/box",
		"Size",
		"Area",
		"Box",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(
		t,
		r,
		"Size satisfies example/pkg/shape.Sizer - renaming it breaks that interface",
	)
	assert.String(t, before, treeOf(t, d))
}

func TestRenameRefusesAMethodSatisfyingAStandardInterface(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface/src"),
	)
	r, e := testService().Rename(
		d,
		"example/pkg/box",
		"String",
		"Label",
		"Box",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "String satisfies fmt.Stringer")
}

func TestRenameRefusesAnInterfaceMethod(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface/src"),
	)
	r, e := testService().Rename(
		d,
		"example/pkg/shape",
		"Size",
		"Area",
		"Sizer",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(
		t,
		r,
		"method or field Size not found on Sizer",
	)
}

func TestRenameAllowsAMethodNoInterfaceNames(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface/src"),
	)
	r, e := testService().Rename(
		d,
		"example/pkg/box",
		"Weight",
		"Mass",
		"Box",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"b.Mass()",
		service_tester.ReadFixtureFile(t, d, "pkg/use/use.go"),
	)
}

func TestRenameRefusesAMethodSatisfyingAReplacingModuleInterface(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	testutil.WriteFile(
		t,
		library,
		"meter/meter.go",
		"package meter\n\ntype Meter struct{}\n\nfunc (m *Meter) Read() int {\n\treturn 1\n}\n",
	)
	testutil.WriteFile(
		t,
		user,
		"pkg/reader/reader.go",
		"package reader\n\nimport \"other.test/lib/meter\"\n\ntype Reader interface {\n\tRead() int\n}\n\nvar Default Reader = &meter.Meter{}\n",
	)
	source := service_tester.ReadFixtureFile(t, library, "meter/meter.go")
	r, e := replacingService(library, user).Rename(
		library,
		"other.test/lib/meter",
		"Read",
		"Measure",
		"Meter",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(
		t,
		r,
		"Read satisfies example/pkg/reader.Reader",
	)
	assert.String(
		t,
		source,
		service_tester.ReadFixtureFile(t, library, "meter/meter.go"),
	)
}
