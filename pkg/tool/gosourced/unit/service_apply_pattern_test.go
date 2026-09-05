package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func gaugeContent(t *testing.T, directory string) string {
	t.Helper()
	b, e := os.ReadFile(filepath.Join(directory, "pkg/gauge/run.go"))
	assert.FatalOnError(t, e)

	return string(b)
}

func TestApplyPatternDryRun(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, apply, e := s.ApplyPattern(
		d,
		"example/pkg/pair",
		"Compare",
		"",
		"func pattern(x int, y int) int {\n\treturn pair.Compare(x, y)\n}",
		"func replacement(x int, y int) int {\n\treturn pair.Compare(y, x) + 1\n}",
		true,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, apply)
	assert.Integer(t, 3, apply.Total)
	assert.Integer(t, 3, apply.Matched)
	assert.Integer(t, 2, apply.Rewritten)
	assert.Integer(t, 1, len(apply.Refused))
	assert.True(t, apply.Applied)
	content := gaugeContent(t, d)
	assert.True(t, !strings.Contains(content, "+ 1"))
}

func TestApplyPatternPartial(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, apply, e := s.ApplyPattern(
		d,
		"example/pkg/pair",
		"Compare",
		"",
		"func pattern(x int, y int) int {\n\treturn pair.Compare(x, y)\n}",
		"func replacement(x int, y int) int {\n\treturn pair.Compare(y, x) + 1\n}",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, apply)
	assert.Integer(t, 2, apply.Rewritten)
	assert.True(t, apply.Applied)
	content := gaugeContent(t, d)
	assert.StringContains(t, "return pair.Compare(n, n) + 1", content)
	assert.StringContains(t, "return pair.Compare(m, n) + 1", content)
	assert.StringContains(t, "// order is deliberate", content)
}

func TestApplyPatternAllOrNothing(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, apply, e := s.ApplyPattern(
		d,
		"example/pkg/pair",
		"Compare",
		"",
		"func pattern(x int, y int) int {\n\treturn pair.Compare(x, y)\n}",
		"func replacement(x int, y int) int {\n\treturn pair.Compare(y, x) + 1\n}",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, apply)
	assert.True(t, !apply.Applied)
	assert.Integer(t, 0, apply.Rewritten)
	assert.True(t, apply.Refusal != "")
	assert.Integer(t, 1, len(apply.Refused))
	content := gaugeContent(t, d)
	assert.True(t, !strings.Contains(content, "+ 1"))
}

func TestApplyPatternImport(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, apply, e := s.ApplyPattern(
		d,
		"fmt",
		"Println",
		"",
		"func pattern(v string) {\n\tfmt.Println(v)\n}",
		"func replacement(v string) {\n\tconsole.Emit(v)\n}",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, apply)
	assert.Integer(t, 2, apply.Rewritten)
	b, f := os.ReadFile(filepath.Join(d, "pkg/monitor/run.go"))
	assert.FatalOnError(t, f)
	content := string(b)
	assert.StringContains(t, "console.Emit(\"ready\")", content)
	assert.StringContains(t, "example/pkg/console", content)
	assert.StringContains(t, "fmt.Println(\"done\", 2)", content)
}

func TestApplyPatternUnknownHole(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("census/src"))
	s := testService()
	r, apply, e := s.ApplyPattern(
		d,
		"example/pkg/pair",
		"Compare",
		"",
		"func pattern(x int, y int) int {\n\treturn pair.Compare(x, y)\n}",
		"func replacement(x int, z int) int {\n\treturn pair.Compare(z, x)\n}",
		true,
		true,
	)
	assert.FatalOnError(t, e)
	assert.True(t, apply == nil)
	testutil.AssertBlockedContains(t, r, "z")
}
