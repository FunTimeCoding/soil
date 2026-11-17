package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestMoveMethodsRetypesReceiver(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-methods/src"))
	r, e := o.Symbol.MoveMethods(
		t.Context(),
		o.Directory,
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
		o.Read("pkg/bravo/greet.go"),
	)
	_, f := os.Stat(filepath.Join(o.Directory, "pkg/alfa/greet.go"))
	assert.True(t, os.IsNotExist(f))
	rewrite_tester.Compiles(t, o.Directory)
}

func TestMoveMethodsQualifiesBackReference(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-methods/src"))
	r, e := o.Symbol.MoveMethods(
		t.Context(),
		o.Directory,
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
	rewrite_tester.FixedContains(t, "prefix → Prefix", r)
	assert.String(
		t,
		"package bravo\n\nimport \"example/pkg/alfa\"\n\nfunc (s *Worker) Label() string {\n\treturn alfa.Prefix() + \"label\"\n}\n",
		o.Read("pkg/bravo/label.go"),
	)
	rewrite_tester.Compiles(t, o.Directory)
}

func TestMoveMethodsKeepsReceiverFieldOfSameType(t *testing.T) {
	o, r := module_tester.MoveMethods(t, "Describe")
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"func (s *Worker) Describe() string {\n\treturn s.Name\n}",
		o.Read("pkg/bravo/describe.go"),
	)
	rewrite_tester.Compiles(t, o.Directory)
}

func TestMoveMethodsRefusesMissingField(t *testing.T) {
	_, r := module_tester.MoveMethods(t, "Count")
	testutil.AssertBlockedContains(
		t,
		r,
		"Count reads s.count, which bravo.Worker lacks",
	)
}

func TestMoveMethodsRefusesFieldOfOtherType(t *testing.T) {
	_, r := module_tester.MoveMethods(t, "Mismatch")
	testutil.AssertBlockedContains(
		t,
		r,
		"Mismatch reads s.Size as int, which bravo.Worker has as string",
	)
}

func TestMoveMethodsRefusesUnmovedSibling(t *testing.T) {
	_, r := module_tester.MoveMethods(t, "Shout")
	testutil.AssertBlockedContains(
		t,
		r,
		"Shout calls s.Loud, which is neither moved nor declared on bravo.Worker",
	)
}

func TestMoveMethodsMovesSiblingsTogether(t *testing.T) {
	o, r := module_tester.MoveMethods(t, "Shout", "Loud")
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.Compiles(t, o.Directory)
}

func TestMoveMethodsKeepsSiblingDeclaredOnTarget(t *testing.T) {
	o, r := module_tester.MoveMethods(t, "Repeat")
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.Compiles(t, o.Directory)
}

func TestMoveMethodsRefusesSiblingOfOtherSignature(t *testing.T) {
	_, r := module_tester.MoveMethods(t, "Mark")
	testutil.AssertBlockedContains(
		t,
		r,
		"Mark calls s.Tag, which is neither moved nor declared on bravo.Worker",
	)
}

func TestMoveMethodsRefusesTakenName(t *testing.T) {
	_, r := module_tester.MoveMethods(t, "Echo")
	testutil.AssertBlockedContains(t, r, "bravo.Worker already has Echo")
}

func TestMoveMethodsRefusesPromotedName(t *testing.T) {
	_, r := module_tester.MoveMethods(t, "Hello")
	testutil.AssertBlockedContains(t, r, "bravo.Worker already has Hello")
}

func TestMoveMethodsRefusesInterfaceMethod(t *testing.T) {
	_, r := module_tester.MoveMethods(t, "String")
	testutil.AssertBlockedContains(
		t,
		r,
		"String satisfies runtime.stringer - moving it off alfa.Service breaks that interface",
	)
}

func TestMoveMethodsRefusesPassedReceiver(t *testing.T) {
	_, r := module_tester.MoveMethods(t, "Pass")
	testutil.AssertBlockedContains(t, r, "Pass uses its receiver s as a value")
}

func TestMoveMethodsRefusesOutsideCaller(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-methods-called/src"),
	)
	r, e := o.Symbol.MoveMethods(
		t.Context(),
		o.Directory,
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
		o.Read("pkg/alfa/greet.go"),
	)
}
