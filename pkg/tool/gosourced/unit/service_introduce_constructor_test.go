package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestIntroduceConstructorRewritesExactSites(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("introduce-constructor/src"),
	)
	r, c, e := testService().IntroduceConstructor(
		d,
		"example/pkg/single",
		"Shape",
		[]string{"Inhale", "Draw"},
		false,
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, "New", c.Name)
	assert.String(t, "pkg/single/new.go", c.File)
	assert.Integer(t, 3, c.Rewritten)
	assert.Integer(t, 5, len(c.Remaining))
	constructor := service_tester.ReadFixtureFile(t, d, "pkg/single/new.go")
	assert.StringContains(
		t,
		"func New(\n\tdraw float64,\n\tinhale float64,\n) *Shape {",
		constructor,
	)
	assert.StringContains(
		t,
		"return &Shape{Draw: draw, Inhale: inhale}",
		constructor,
	)
	shapes := service_tester.ReadFixtureFile(t, d, "pkg/caller/shapes.go")
	assert.StringContains(t, "a := single.New(1, 2)", shapes)
	assert.StringContains(t, "return single.New(1, 2)\n}", shapes)
	assert.StringContains(t, "c := single.New(1, 2)\n\tc.Hold = 3", shapes)
	assert.StringContains(
		t,
		"[]*single.Shape{{Draw: 1, Inhale: 2, Wait: time.Second}}",
		shapes,
	)
	assert.StringContains(t, "return &single.Shape{Draw: 1}", shapes)
	assert.StringContains(t, "return single.Shape{Draw: 1, Inhale: 2}", shapes)
	assert.StringContains(
		t,
		"return &single.Shape{Inhale: two(), Draw: one()}",
		shapes,
	)
	assert.StringContains(t, "return new(single.Shape)", shapes)
}

func TestIntroduceConstructorReportsWhyASiteStays(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("introduce-constructor/src"),
	)
	_, c, e := testService().IntroduceConstructor(
		d,
		"example/pkg/single",
		"Shape",
		[]string{"Draw", "Inhale"},
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, c.Rewritten)
	assert.Strings(
		t,
		[]string{
			"&{Draw, Inhale, Hold} - sets Hold beyond the constructor",
			"&{Draw, Inhale} nested - evaluation order",
			"&{Draw} nested - misses Inhale",
			"new() nested - misses Draw, Inhale",
			"{Draw, Inhale, Wait} elided nested - sets Wait beyond the constructor",
			"{Draw, Inhale} value nested - value literal",
		},
		remainingShapes(c),
	)
}

func TestIntroduceConstructorImportsParameterTypes(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("introduce-constructor/src"),
	)
	_, _, e := testService().IntroduceConstructor(
		d,
		"example/pkg/single",
		"Shape",
		[]string{"Wait"},
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	constructor := service_tester.ReadFixtureFile(t, d, "pkg/single/new.go")
	assert.StringContains(t, "import \"time\"", constructor)
	assert.StringContains(
		t,
		"func New(wait time.Duration) *Shape {",
		constructor,
	)
}

func TestIntroduceConstructorNamesBagConstructorsAfterTheType(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("introduce-constructor/src"),
	)
	_, c, e := testService().IntroduceConstructor(
		d,
		"example/pkg/bag",
		"Payload",
		nil,
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "NewPayload", c.Name)
	assert.String(t, "pkg/bag/new_payload.go", c.File)
	constructor := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/bag/new_payload.go",
	)
	assert.StringContains(t, "func NewPayload() *Payload {", constructor)
	payloads := service_tester.ReadFixtureFile(t, d, "pkg/caller/payloads.go")
	assert.StringContains(t, "p := bag.NewPayload()", payloads)
}

func TestIntroduceConstructorTakesParametersFromSites(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("introduce-constructor/src"),
	)
	_, c, e := testService().IntroduceConstructor(
		d,
		"example/pkg/bag",
		"Header",
		nil,
		true,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"Key"}, c.Parameters)
	assert.Integer(t, 1, c.Rewritten)
	payloads := service_tester.ReadFixtureFile(t, d, "pkg/caller/payloads.go")
	assert.StringContains(t, "bag.NewHeader(\"charlie\")", payloads)
	assert.StringContains(
		t,
		"&bag.Header{Key: \"alfa\", Value: \"bravo\"}",
		payloads,
	)
}

func TestIntroduceConstructorRefusesATaintedPackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("introduce-constructor/src"),
	)
	r, c, e := testService().IntroduceConstructor(
		d,
		"example/pkg/tainted",
		"Record",
		nil,
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	assert.Nil(t, c)
	testutil.AssertBlockedContains(
		t,
		r,
		"Record shares example/pkg/tainted with receiver struct Service - move it out first with extract_type",
	)
}

func TestIntroduceConstructorRefusesAnUnknownParameter(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("introduce-constructor/src"),
	)
	r, c, e := testService().IntroduceConstructor(
		d,
		"example/pkg/single",
		"Shape",
		[]string{"Missing"},
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	assert.Nil(t, c)
	testutil.AssertBlockedContains(t, r, "no field Missing")
}

func TestIntroduceConstructorDryRunWritesNothing(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("introduce-constructor/src"),
	)
	_, c, e := testService().IntroduceConstructor(
		d,
		"example/pkg/single",
		"Shape",
		[]string{"Draw", "Inhale"},
		false,
		false,
		true,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, c.Rewritten)
	_, f := os.Stat(filepath.Join(d, "pkg/single/new.go"))
	assert.True(t, os.IsNotExist(f))
	shapes := service_tester.ReadFixtureFile(t, d, "pkg/caller/shapes.go")
	assert.StringContains(t, "a := &single.Shape{Draw: 1, Inhale: 2}", shapes)
}
