package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveMethodsRoutesCallersThroughHolder(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods-held/src"),
	)
	r := moveHeldGreet(t, d)
	testutil.AssertBlocked(t, r, 0)
	assertFixedContains(t, "Greet → Worker.Greet", r)
	assert.StringContains(
		t,
		"return s.Worker.Greet(s.Name)",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/hello.go"),
	)
	assert.StringContains(
		t,
		"return s.Worker.Greet(\"charlie\")",
		service_tester.ReadFixtureFile(t, d, "pkg/charlie/use.go"),
	)
	assert.StringContains(
		t,
		"return s.Worker.Greet(\"value\")",
		service_tester.ReadFixtureFile(t, d, "pkg/charlie/value.go"),
	)
	assert.StringContains(
		t,
		"return s.Worker.Greet\n",
		service_tester.ReadFixtureFile(t, d, "pkg/charlie/hold.go"),
	)
	assertCompiles(t, d)
}

func TestMoveMethodsRefusesTwoHolders(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods-held/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"pkg/alfa/service.go",
		"package alfa\n\nimport \"example/pkg/bravo\"\n\ntype Service struct {\n\tName   string\n\tWorker *bravo.Worker\n\tSpare  *bravo.Worker\n}\n",
	)
	testutil.AssertBlockedContains(
		t,
		moveHeldGreet(t, d),
		"alfa.Service holds no single field of type bravo.Worker",
	)
}

func TestMoveMethodsRefusesUnexportedHolderAcrossPackages(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods-held/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"pkg/alfa/service.go",
		"package alfa\n\nimport \"example/pkg/bravo\"\n\ntype Service struct {\n\tName   string\n\tworker *bravo.Worker\n}\n",
	)
	testutil.AssertBlockedContains(
		t,
		moveHeldGreet(t, d),
		"where the unexported worker cannot be reached",
	)
}

func TestMoveMethodsRefusesCallThroughEmbedding(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods-held/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"pkg/charlie/outer.go",
		"package charlie\n\nimport \"example/pkg/alfa\"\n\ntype Outer struct {\n\t*alfa.Service\n}\n\nfunc (o Outer) Run() string {\n\treturn o.Greet(\"outer\")\n}\n",
	)
	testutil.AssertBlockedContains(
		t,
		moveHeldGreet(t, d),
		"through an embedding",
	)
}

func TestMoveMethodsRefusesValueHolderOnValue(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods-held/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"pkg/alfa/service.go",
		"package alfa\n\nimport \"example/pkg/bravo\"\n\ntype Service struct {\n\tName   string\n\tWorker bravo.Worker\n}\n",
	)
	testutil.AssertBlockedContains(
		t,
		moveHeldGreet(t, d),
		"on a value, where alfa.Service.Worker may not be addressable",
	)
}

func TestMoveMethodsRefusesMethodExpression(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods-held/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"pkg/charlie/expression.go",
		"package charlie\n\nimport \"example/pkg/alfa\"\n\nvar Greeting = (*alfa.Service).Greet\n",
	)
	testutil.AssertBlockedContains(
		t,
		moveHeldGreet(t, d),
		"as a method expression",
	)
}
