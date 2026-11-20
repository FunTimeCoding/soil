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

func TestMoveMethodsRoutesCallersThroughHolder(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-methods-held/src"),
	)
	r := module_tester.New(t).MoveHeldGreet(o)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.New(t).FixedContains("Greet → Worker.Greet", r)
	assert.StringContains(
		t,
		"return s.Worker.Greet(s.Name)",
		o.Read("pkg/alfa/hello.go"),
	)
	assert.StringContains(
		t,
		"return s.Worker.Greet(\"charlie\")",
		o.Read("pkg/charlie/use.go"),
	)
	assert.StringContains(
		t,
		"return s.Worker.Greet(\"value\")",
		o.Read("pkg/charlie/value.go"),
	)
	assert.StringContains(
		t,
		"return s.Worker.Greet\n",
		o.Read("pkg/charlie/hold.go"),
	)
	rewrite_tester.New(t).Compiles(o.Directory)
}

func TestMoveMethodsRefusesTwoHolders(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-methods-held/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"pkg/alfa/service.go",
		"package alfa\n\nimport \"example/pkg/bravo\"\n\ntype Service struct {\n\tName   string\n\tWorker *bravo.Worker\n\tSpare  *bravo.Worker\n}\n",
	)
	testutil.AssertBlockedContains(
		t,
		module_tester.New(t).MoveHeldGreet(o),
		"alfa.Service holds no single field of type bravo.Worker",
	)
}

func TestMoveMethodsRefusesUnexportedHolderAcrossPackages(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-methods-held/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"pkg/alfa/service.go",
		"package alfa\n\nimport \"example/pkg/bravo\"\n\ntype Service struct {\n\tName   string\n\tworker *bravo.Worker\n}\n",
	)
	testutil.AssertBlockedContains(
		t,
		module_tester.New(t).MoveHeldGreet(o),
		"where the unexported worker cannot be reached",
	)
}

func TestMoveMethodsRefusesCallThroughEmbedding(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-methods-held/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"pkg/charlie/outer.go",
		"package charlie\n\nimport \"example/pkg/alfa\"\n\ntype Outer struct {\n\t*alfa.Service\n}\n\nfunc (o Outer) Run() string {\n\treturn o.Greet(\"outer\")\n}\n",
	)
	testutil.AssertBlockedContains(
		t,
		module_tester.New(t).MoveHeldGreet(o),
		"through an embedding",
	)
}

func TestMoveMethodsRefusesValueHolderOnValue(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-methods-held/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"pkg/alfa/service.go",
		"package alfa\n\nimport \"example/pkg/bravo\"\n\ntype Service struct {\n\tName   string\n\tWorker bravo.Worker\n}\n",
	)
	testutil.AssertBlockedContains(
		t,
		module_tester.New(t).MoveHeldGreet(o),
		"on a value, where alfa.Service.Worker may not be addressable",
	)
}

func TestMoveMethodsRefusesMethodExpression(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-methods-held/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"pkg/charlie/expression.go",
		"package charlie\n\nimport \"example/pkg/alfa\"\n\nvar Greeting = (*alfa.Service).Greet\n",
	)
	testutil.AssertBlockedContains(
		t,
		module_tester.New(t).MoveHeldGreet(o),
		"as a method expression",
	)
}
