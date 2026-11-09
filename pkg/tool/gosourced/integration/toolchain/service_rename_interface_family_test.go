package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameOfInterfaceMethodRenamesItsFamily(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface-family/src"),
	)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "pkg/box/box.go"),
	)
	assert.StringContains(
		t,
		"func (m *Mock) Area() int",
		service_tester.ReadFixtureFile(t, d, "pkg/mock/mock.go"),
	)
	assert.StringContains(
		t,
		"Area() int",
		service_tester.ReadFixtureFile(t, d, "pkg/shape/measured.go"),
	)
	assert.StringContains(
		t,
		"return s.Area() + m.Area() + b.Area()",
		service_tester.ReadFixtureFile(t, d, "pkg/use/use.go"),
	)
}

func TestRenameOfInterfaceMethodListsFamilyBeforeWriting(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface-family/src"),
	)
	before := directoryTree(t, d)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/shape",
		"Size",
		"Area",
		"Sizer",
		true,
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		[]string{
			"func (*example/pkg/box.Box).Size() int",
			"func (*example/pkg/mock.Mock).Size() int",
			"func (example/pkg/shape.Measured).Size() int",
			"func (example/pkg/shape.Sizer).Size() int",
		},
		familyMembers(r),
	)
	assert.String(t, before, directoryTree(t, d))
}

func TestRenameOfInterfaceMethodRefusesStandardLibraryFamily(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface-family/src"),
	)
	before := directoryTree(t, d)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/shape",
		"String",
		"Label",
		"Labeler",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(
		t,
		r,
		"example/pkg/box.Box satisfies fmt.Stringer outside the workspace",
	)
	assert.String(t, before, directoryTree(t, d))
}

func TestRenameOfInterfaceMethodRefusesCollisionAnywhereInFamily(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface-family/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"pkg/mock/area.go",
		"package mock\n\nfunc (m *Mock) Area() int {\n\treturn 3\n}\n",
	)
	before := directoryTree(t, d)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/shape",
		"Size",
		"Area",
		"Sizer",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(
		t,
		r,
		"example/pkg/mock.Mock already has Area",
	)
	assert.String(t, before, directoryTree(t, d))
}

func TestRenameOfEmbeddedInterfaceMethodNamesItsDeclarer(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface-family/src"),
	)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/shape",
		"Size",
		"Area",
		"Solid",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(
		t,
		r,
		"Size is declared in example/pkg/shape.Sizer, embedded in Solid - rename it there",
	)
}

func TestRenameOfImplementationPointsToItsInterface(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface-family/src"),
	)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/box",
		"Size",
		"Area",
		"Box",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(
		t,
		r,
		"Size satisfies example/pkg/shape.Sizer - renaming it breaks that interface; rename Sizer.Size to rename them together",
	)
}

func TestRenameOfInterfaceMethodRenamesWideStructuralFamily(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-interface-wide/src"),
	)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/shape",
		"Size",
		"Area",
		"Sizer",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.Count(t, 6, familyMembers(r))
	assert.StringContains(
		t,
		"return (&Alfa{}).Area() + Bravo{}.Area() + d.Area() + Echo(4).Area()",
		service_tester.ReadFixtureFile(t, d, "pkg/part/foxtrot.go"),
	)
	assert.StringContains(
		t,
		"func (f *Foxtrot) Area() int",
		service_tester.ReadFixtureFile(t, d, "pkg/part/foxtrot.go"),
	)
}
