package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestApplyPatternDryRun(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("pattern/src"))
	r, apply, e := o.Pattern.ApplyPattern(
		t.Context(),
		o.Directory,
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
	content := service_tester.GaugeContent(t, o.Directory)
	assert.True(t, !strings.Contains(content, "+ 1"))
}

func TestApplyPatternPartial(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("pattern/src"))
	r, apply, e := o.Pattern.ApplyPattern(
		t.Context(),
		o.Directory,
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
	content := service_tester.GaugeContent(t, o.Directory)
	assert.StringContains(t, "return pair.Compare(n, n) + 1", content)
	assert.StringContains(t, "return pair.Compare(m, n) + 1", content)
	assert.StringContains(t, "// order is deliberate", content)
}

func TestApplyPatternAllOrNothing(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("pattern/src"))
	r, apply, e := o.Pattern.ApplyPattern(
		t.Context(),
		o.Directory,
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
	content := service_tester.GaugeContent(t, o.Directory)
	assert.True(t, !strings.Contains(content, "+ 1"))
}

func TestApplyPatternImport(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("pattern/src"))
	r, apply, e := o.Pattern.ApplyPattern(
		t.Context(),
		o.Directory,
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
	b, f := os.ReadFile(filepath.Join(o.Directory, "pkg/monitor/run.go"))
	assert.FatalOnError(t, f)
	content := string(b)
	assert.StringContains(t, "console.Emit(\"ready\")", content)
	assert.StringContains(t, "example/pkg/console", content)
	assert.StringContains(t, "fmt.Println(\"done\", 2)", content)
}
