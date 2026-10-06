package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestFindLiteralsGroupsOutsideSitesByShape(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("find-literals/src"),
	)
	r, literals, e := testService().FindLiterals(
		d,
		"example/pkg/target",
		"Shape",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, "Shape", literals.Type)
	assert.Integer(t, 10, literals.Total)
	assert.Integer(t, 1, literals.Inside)
	assert.Integer(t, 1, literals.Expected)
	assert.Integer(t, 7, len(literals.Groups))
	assert.String(t, "&{Draw, Inhale}", literals.Groups[0].Shape)
	assert.Integer(t, 2, len(literals.Groups[0].Locations))
	assert.String(t, "pkg/caller/run.go", literals.Groups[0].Locations[0].File)
	assert.Integer(t, 9, literals.Groups[0].Locations[0].Line)
	assert.String(
		t,
		"&target.Shape{Draw: 1, Inhale: 2}",
		literals.Groups[0].Exemplar,
	)
	assert.String(t, "&{Draw, Inhale, Hold}", literals.Groups[1].Shape)
	assert.String(
		t,
		"&{Draw, Strength, Inhale, Hold} unkeyed",
		literals.Groups[2].Shape,
	)
	assert.String(t, "&{Hold} nested", literals.Groups[3].Shape)
	assert.String(t, "new() nested", literals.Groups[4].Shape)
	assert.String(t, "{Draw} elided nested", literals.Groups[5].Shape)
	assert.String(t, "{} value", literals.Groups[6].Shape)
}

func TestFindLiteralsCollapsesAMultilineExemplar(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("find-literals-exemplar/src"),
	)
	_, literals, e := testService().FindLiterals(
		d,
		"example/pkg/target",
		"Shape",
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"&target.Shape{Draw: 1, Label: fmt.Sprintf(\"%s-%d\", \"alfa\", 2)}",
		literals.Groups[0].Exemplar,
	)
}

func TestFindLiteralsRefusesAMissingType(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("find-literals/src"),
	)
	r, literals, e := testService().FindLiterals(
		d,
		"example/pkg/target",
		"Missing",
	)
	assert.FatalOnError(t, e)
	assert.Nil(t, literals)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestFindLiteralsRefusesANonStruct(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("find-literals/src"),
	)
	r, literals, e := testService().FindLiterals(
		d,
		"example/pkg/target",
		"DefaultShape",
	)
	assert.FatalOnError(t, e)
	assert.Nil(t, literals)
	testutil.AssertBlockedContains(t, r, "not a struct type")
}
