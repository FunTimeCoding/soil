package redundant_type_name

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/redundant_type_name"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestFlagsBagStructRepeatingPackageName(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/flagged")
	redundant_type_name.Check(p, results)
	testutil.AssertBlocked(t, results, 2)
	testutil.AssertLineBlocked(t, results, "types.go", 3)
	testutil.AssertLineBlocked(t, results, "types.go", 7)
	testutil.AssertBlockedContains(
		t,
		results,
		"struct LinkResponse repeats the bag's name response - rename it Link",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"struct ResponseTag repeats the bag's name response - rename it Tag",
	)
}

func TestTakenNameAsksForDistinction(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/collision")
	redundant_type_name.Check(p, results)
	testutil.AssertBlocked(t, results, 1)
	testutil.AssertBlockedContains(
		t,
		results,
		"struct LinkResponse repeats the bag's name response, and Link is taken - name what tells LinkResponse from Link",
	)
}

func TestOwnPackageIsClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/own")
	redundant_type_name.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestReceiverPackageIsClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/receiver")
	redundant_type_name.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestLoneStructIsClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/lone")
	redundant_type_name.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestGeneratedStructIsIgnored(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/generated")
	redundant_type_name.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
