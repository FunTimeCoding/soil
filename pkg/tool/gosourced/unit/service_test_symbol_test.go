package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestRenameTestFunctionInTestOnlyPackage(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("test-symbols/src"))
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/flow",
		"TestFlowStart",
		"TestFlowBegin",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	renamed := readFixtureFile(t, d, "pkg/flow/flow_test.go")
	assert.StringContains(t, "func TestFlowBegin(", renamed)
}

func TestRenameTestFunctionBesideProduction(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("test-symbols/src"))
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target",
		"TestCompute",
		"TestComputeValue",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	renamed := readFixtureFile(t, d, "pkg/target/compute_test.go")
	assert.StringContains(t, "func TestComputeValue(", renamed)
}

func TestListCallsCountsTestVariantPackageOnce(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("test-symbols/src"))
	s := testService()
	r, inventory, e := s.ListCalls(d, "example/pkg/measure", 0)
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
	d := testutil.PrepareTestPackage(t, serviceTestdata("test-symbols/src"))
	s := testService()
	r, e := s.Rename(d, "example/pkg/target", "Compute", "Calculate", "", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	declaration := readFixtureFile(t, d, "pkg/target/compute.go")
	assert.StringContains(t, "func Calculate(", declaration)
	inTest := readFixtureFile(t, d, "pkg/target/compute_test.go")
	assert.StringContains(t, "Calculate()", inTest)
	crossPackage := readFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "target.Calculate()", crossPackage)
}
