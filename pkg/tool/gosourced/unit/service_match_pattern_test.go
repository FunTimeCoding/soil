package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestMatchPattern(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, match, e := s.MatchPattern(
		d,
		"example/pkg/client",
		"Ready",
		"Client",
		"func pattern(c *client.Client) {\n\tif c.Ready() {\n\t}\n}",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, match)
	assert.String(t, "Ready", match.Symbol)
	assert.Integer(t, 4, match.Total)
	assert.Integer(t, 3, match.Matched)
	assert.Integer(t, 1, len(match.Unmatched))
	assert.String(t, "if c.Ready() && IDENT < INT {", match.Unmatched[0].Shape)
	assert.String(
		t,
		"if c.Ready() && retries < 3 {",
		match.Unmatched[0].Exemplar,
	)
	assert.Integer(t, 1, len(match.Unmatched[0].Locations))
	assert.String(t, "pkg/web/health.go", match.Unmatched[0].Locations[0].File)
	assert.Integer(t, 22, match.Unmatched[0].Locations[0].Line)
}

func TestMatchPatternImported(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, match, e := s.MatchPattern(
		d,
		"fmt",
		"Println",
		"",
		"func pattern(v string) {\n\tfmt.Println(v)\n}",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, match)
	assert.Integer(t, 4, match.Total)
	assert.Integer(t, 2, match.Matched)
	assert.Integer(t, 2, len(match.Unmatched))
	assert.String(t, "fmt.Println(\"done\", 2)", match.Unmatched[0].Exemplar)
	assert.String(t, "pkg/monitor/run.go", match.Unmatched[0].Locations[0].File)
	assert.String(t, "fmt.Println(7)", match.Unmatched[1].Exemplar)
	assert.String(t, "pkg/gauge/run.go", match.Unmatched[1].Locations[0].File)
}

func TestMatchPatternHoleConsistency(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, match, e := s.MatchPattern(
		d,
		"example/pkg/pair",
		"Compare",
		"",
		"func pattern(x int) int {\n\treturn pair.Compare(x, x)\n}",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, match)
	assert.Integer(t, 3, match.Total)
	assert.Integer(t, 1, match.Matched)
	assert.Integer(t, 2, len(match.Unmatched))
	assert.String(t, "return pair.Compare(m, n)", match.Unmatched[0].Exemplar)
	assert.String(t, "return pair.Compare(n, m)", match.Unmatched[1].Exemplar)
}

func TestMatchPatternType(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, match, e := s.MatchPattern(
		d,
		"example/pkg/client",
		"Client",
		"",
		"func pattern(c *client.Client, x []any) {\n\tif c.Client(x...) {\n\t}\n}",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, match)
	assert.Integer(t, 5, match.Total)
	assert.Integer(t, 4, match.Matched)
	assert.Integer(t, 1, len(match.Unmatched))
	assert.String(
		t,
		"if c.Ready() && retries < 3 {",
		match.Unmatched[0].Exemplar,
	)
}

func TestMatchPatternSpreadHoleType(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, match, e := s.MatchPattern(
		d,
		"example/pkg/client",
		"Client",
		"",
		"func pattern(c *client.Client, x []int) {\n\tif c.Client(x...) {\n\t}\n}",
	)
	assert.FatalOnError(t, e)
	assert.True(t, match == nil)
	testutil.AssertBlockedContains(t, r, "[]any")
}

func TestListCalls(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, inventory, e := s.ListCalls(d, "example/pkg/monitor", 0)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, inventory)
	assert.Integer(t, 3, inventory.Total)
	assert.Integer(t, 0, inventory.More)
	assert.String(t, "fmt.Println", inventory.Calls[0].Name)
	assert.Integer(t, 3, inventory.Calls[0].Count)
	assert.String(
		t,
		"(*example/pkg/client.Client).Ready",
		inventory.Calls[1].Name,
	)
	assert.Integer(t, 1, inventory.Calls[1].Count)
	assert.String(
		t,
		"(*example/pkg/client.Client).Steady",
		inventory.Calls[2].Name,
	)
}

func TestListCallsUnknownRegion(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, inventory, e := s.ListCalls(d, "example/pkg/missing", 0)
	assert.FatalOnError(t, e)
	assert.True(t, inventory == nil)
	testutil.AssertBlockedContains(t, r, "missing")
}

func TestMatchPatternWhitespace(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, match, e := s.MatchPattern(
		d,
		"example/pkg/client",
		"Ready",
		"Client",
		"func   pattern( c  *client.Client )  {\n\n\tif   c . Ready( )   {\n\t}\n\n}",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, match)
	assert.Integer(t, 4, match.Total)
	assert.Integer(t, 3, match.Matched)
	assert.Integer(t, 1, len(match.Unmatched))
}
