package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestMoveMethodsRetypesReceiver(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods/src"),
	)
	r, e := service_tester.NewService(t).Symbol.MoveMethods(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Service",
		[]string{"Greet"},
		"example/pkg/bravo",
		"Worker",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package bravo\n\nimport \"strings\"\n\nfunc (s *Worker) Greet(who string) string {\n\treturn strings.ToUpper(who)\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/greet.go"),
	)
	_, f := os.Stat(filepath.Join(d, "pkg/alfa/greet.go"))
	assert.True(t, os.IsNotExist(f))
	assertCompiles(t, d)
}

func TestMoveMethodsQualifiesBackReference(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods/src"),
	)
	r, e := service_tester.NewService(t).Symbol.MoveMethods(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Service",
		[]string{"Label"},
		"example/pkg/bravo",
		"Worker",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assertFixedContains(t, "prefix → Prefix", r)
	assert.String(
		t,
		"package bravo\n\nimport \"example/pkg/alfa\"\n\nfunc (s *Worker) Label() string {\n\treturn alfa.Prefix() + \"label\"\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/label.go"),
	)
	assertCompiles(t, d)
}

func TestMoveMethodsKeepsReceiverFieldOfSameType(t *testing.T) {
	d, r := moveFixtureMethods(t, "Describe")
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"func (s *Worker) Describe() string {\n\treturn s.Name\n}",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/describe.go"),
	)
	assertCompiles(t, d)
}

func TestMoveMethodsRefusesMissingField(t *testing.T) {
	_, r := moveFixtureMethods(t, "Count")
	testutil.AssertBlockedContains(
		t,
		r,
		"Count reads s.count, which bravo.Worker lacks",
	)
}

func TestMoveMethodsRefusesFieldOfOtherType(t *testing.T) {
	_, r := moveFixtureMethods(t, "Mismatch")
	testutil.AssertBlockedContains(
		t,
		r,
		"Mismatch reads s.Size as int, which bravo.Worker has as string",
	)
}

func TestMoveMethodsRefusesUnmovedSibling(t *testing.T) {
	_, r := moveFixtureMethods(t, "Shout")
	testutil.AssertBlockedContains(
		t,
		r,
		"Shout calls s.Loud, which is neither moved nor declared on bravo.Worker",
	)
}

func TestMoveMethodsMovesSiblingsTogether(t *testing.T) {
	d, r := moveFixtureMethods(t, "Shout", "Loud")
	testutil.AssertBlocked(t, r, 0)
	assertCompiles(t, d)
}

func TestMoveMethodsKeepsSiblingDeclaredOnTarget(t *testing.T) {
	d, r := moveFixtureMethods(t, "Repeat")
	testutil.AssertBlocked(t, r, 0)
	assertCompiles(t, d)
}

func TestMoveMethodsRefusesSiblingOfOtherSignature(t *testing.T) {
	_, r := moveFixtureMethods(t, "Mark")
	testutil.AssertBlockedContains(
		t,
		r,
		"Mark calls s.Tag, which is neither moved nor declared on bravo.Worker",
	)
}

func TestMoveMethodsRefusesTakenName(t *testing.T) {
	_, r := moveFixtureMethods(t, "Echo")
	testutil.AssertBlockedContains(t, r, "bravo.Worker already has Echo")
}

func TestMoveMethodsRefusesPromotedName(t *testing.T) {
	_, r := moveFixtureMethods(t, "Hello")
	testutil.AssertBlockedContains(t, r, "bravo.Worker already has Hello")
}

func TestMoveMethodsRefusesInterfaceMethod(t *testing.T) {
	_, r := moveFixtureMethods(t, "String")
	testutil.AssertBlockedContains(
		t,
		r,
		"String satisfies runtime.stringer - moving it off alfa.Service breaks that interface",
	)
}

func TestMoveMethodsRefusesPassedReceiver(t *testing.T) {
	_, r := moveFixtureMethods(t, "Pass")
	testutil.AssertBlockedContains(t, r, "Pass uses its receiver s as a value")
}

func TestMoveMethodsRefusesOutsideCaller(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods-called/src"),
	)
	r, e := service_tester.NewService(t).Symbol.MoveMethods(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Service",
		[]string{"Greet"},
		"example/pkg/bravo",
		"Worker",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "Greet is called at")
	assert.StringContains(
		t,
		"func (s *Service) Greet",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/greet.go"),
	)
}
