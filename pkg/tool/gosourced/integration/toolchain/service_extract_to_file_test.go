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

func TestExtractFunction(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-function/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"FormatName",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := o.Read("pkg/target/format_name.go")
	assert.StringContains(t, "func FormatName(", extracted)
	assert.StringContains(t, "fmt.Sprintf", extracted)
	assert.StringContains(t, "// FormatName renders a labeled name.", extracted)
	source := o.Read("pkg/target/combined.go")
	assert.True(t, !strings.Contains(source, "func FormatName("))
	assert.True(t, !strings.Contains(source, "// FormatName renders"))
	assert.StringContains(t, "func TrimName(", source)
	assert.StringContains(t, "func PlainName(", source)
}

func TestExtractFunctionRemovesUnusedImport(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-function/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"TrimName",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := o.Read("pkg/target/trim_name.go")
	assert.StringContains(t, "func TrimName(", extracted)
	assert.StringContains(t, "strings", extracted)
	source := o.Read("pkg/target/combined.go")
	assert.True(t, !strings.Contains(source, "strings"))
	assert.StringContains(t, "fmt", source)
}

func TestExtractFunctionNoImports(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-function/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"PlainName",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := o.Read("pkg/target/plain_name.go")
	assert.StringContains(t, "func PlainName(", extracted)
	assert.True(t, !strings.Contains(extracted, "import"))
}

func TestExtractRenamesSourceWhenOneRemains(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-last-pair/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"FormatName",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	extracted := o.Read("pkg/target/format_name.go")
	assert.StringContains(t, "func FormatName(", extracted)
	renamed := o.Read("pkg/target/summarize_name.go")
	assert.StringContains(t, "func SummarizeName(", renamed)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/target/combined.go"))
	assert.True(t, os.IsNotExist(e))
}

func TestExtractTypeToFile(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-identity/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"Widget",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := o.Read("pkg/target/widget.go")
	assert.StringContains(t, "type Widget struct", extracted)
	assert.StringContains(t, "// Widget holds a display name.", extracted)
	source := o.Read("pkg/target/combined.go")
	assert.True(t, !strings.Contains(source, "type Widget struct"))
	assert.StringContains(t, "func (w *Widget) Trimmed(", source)
	assert.StringContains(t, "func NewWidget(", source)
}

func TestExtractLeavesTypeWithoutRename(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-identity/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"NewWidget",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/target/combined.go")
	assert.StringContains(t, "type Widget struct", source)
	assert.StringContains(t, "func (w *Widget) Trimmed(", source)
}

func TestExtractFromTestFileKeepsItsName(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-test-file/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/target_test.go",
		"newWidget",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/target/target_test.go")
	assert.StringContains(t, "func TestWidgetName(", source)
	assert.True(t, !strings.Contains(source, "func newWidget("))
}

func TestExtractFromTestFileCarriesRenamedImport(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-test-file/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/target_test.go",
		"newWidget",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := o.Read("pkg/target/new_widget.go")
	assert.StringContains(t, "func newWidget(", extracted)
	assert.StringContains(t, "example/pkg/go-widget", extracted)
}

func TestExtractFromTestFileCarriesBuildConstraint(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-test-file/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/target_test.go",
		"newWidget",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := o.Read("pkg/target/new_widget.go")
	assert.Prefix(t, "//go:build ci", extracted)
}

func TestExtractResolvesTagGatedDependency(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-tagged-dependency/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"Describe",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := o.Read("pkg/target/describe.go")
	assert.StringContains(t, "func Describe(", extracted)
	assert.StringContains(t, "example/pkg/gated", extracted)
}
