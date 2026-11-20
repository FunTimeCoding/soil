package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameRefusesMethodSatisfyingInterface(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-interface/src"),
	)
	before := rewrite_tester.New(t).DirectoryTree(o.Directory)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
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
	assert.String(t, before, rewrite_tester.New(t).DirectoryTree(o.Directory))
}

func TestRenameRefusesMethodSatisfyingStandardInterface(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-interface/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/box",
		"String",
		"Label",
		"Box",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "String satisfies fmt.Stringer")
}

func TestRenameOfInterfaceMethodRenamesItsImplementation(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-interface/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/shape",
		"Size",
		"Area",
		"Sizer",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"func (b *Box) Area() int",
		o.Read("pkg/box/box.go"),
	)
	assert.StringContains(t, "s.Area()", o.Read("pkg/use/use.go"))
}

func TestRenameAllowsMethodNoInterfaceNames(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-interface/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/box",
		"Weight",
		"Mass",
		"Box",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(t, "b.Mass()", o.Read("pkg/use/use.go"))
}

func TestRenameRefusesMethodSatisfyingReplacingModuleInterface(t *testing.T) {
	library, o := module_tester.New(t).ReplacedLibrary("rename-replacer")
	system_tester.New(t).Write(
		library,
		"meter/meter.go",
		"package meter\n\ntype Meter struct{}\n\nfunc (m *Meter) Read() int {\n\treturn 1\n}\n",
	)
	system_tester.New(t).Write(
		o.Directory,
		"pkg/reader/reader.go",
		"package reader\n\nimport \"other.test/lib/meter\"\n\ntype Reader interface {\n\tRead() int\n}\n\nvar Default Reader = &meter.Meter{}\n",
	)
	source := system_tester.New(t).Read(library, "meter/meter.go")
	r, e := o.Symbol.Rename(
		t.Context(),
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
		system_tester.New(t).Read(library, "meter/meter.go"),
	)
}
