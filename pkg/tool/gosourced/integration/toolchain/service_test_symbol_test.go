package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameTestFunctionInTestOnlyPackage(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("test-symbols/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/flow",
		"TestFlowStart",
		"TestFlowBegin",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	renamed := o.Read("pkg/flow/flow_test.go")
	assert.StringContains(t, "func TestFlowBegin(", renamed)
}

func TestRenameTestFunctionBesideProduction(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("test-symbols/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"TestCompute",
		"TestComputeValue",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	renamed := o.Read("pkg/target/compute_test.go")
	assert.StringContains(t, "func TestComputeValue(", renamed)
}

func TestListCallsCountsTestVariantPackageOnce(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("test-symbols/src"))
	r, inventory, e := o.Reference.ListCalls(
		o.Directory,
		"example/pkg/measure",
		0,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	count := 0

	for _, c := range inventory.Calls {
		if c.Name == "example/pkg/helper.Base" {
			count = c.Count
		}
	}

	assert.Integer(t, 1, count)
}

func TestRenameProductionSymbolRewritesTestFile(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("test-symbols/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Compute",
		"Calculate",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	declaration := o.Read("pkg/target/calculate.go")
	assert.StringContains(t, "func Calculate(", declaration)
	testFile := o.Read("pkg/target/compute_test.go")
	assert.StringContains(t, "Calculate()", testFile)
	crossPackage := o.Read("pkg/caller/run.go")
	assert.StringContains(t, "target.Calculate()", crossPackage)
}
