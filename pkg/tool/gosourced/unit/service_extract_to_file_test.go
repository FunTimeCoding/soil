package unit

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
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-function/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "FormatName", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/format_name.go",
	)
	assert.StringContains(t, "func FormatName(", extracted)
	assert.StringContains(t, "fmt.Sprintf", extracted)
	assert.StringContains(t, "// FormatName renders a labeled name.", extracted)
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/combined.go")
	assert.True(t, !strings.Contains(source, "func FormatName("))
	assert.True(t, !strings.Contains(source, "// FormatName renders"))
	assert.StringContains(t, "func TrimName(", source)
	assert.StringContains(t, "func PlainName(", source)
}

func TestExtractFunctionRemovesUnusedImport(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-function/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "TrimName", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := service_tester.ReadFixtureFile(t, d, "pkg/target/trim_name.go")
	assert.StringContains(t, "func TrimName(", extracted)
	assert.StringContains(t, "strings", extracted)
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/combined.go")
	assert.True(t, !strings.Contains(source, "strings"))
	assert.StringContains(t, "fmt", source)
}

func TestExtractFunctionNoImports(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-function/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "PlainName", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/plain_name.go",
	)
	assert.StringContains(t, "func PlainName(", extracted)
	assert.True(t, !strings.Contains(extracted, "import"))
}

func TestExtractFunctionNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-function/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "Missing", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestExtractFunctionFileNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-function/src"),
	)
	s := testService()
	_, e := s.ExtractToFile(d, "pkg/target/missing.go", "Something", false)
	assert.True(t, e != nil)
}

func TestExtractFunctionTargetExists(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-function/src"),
	)
	target := filepath.Join(d, "pkg/target/format_name.go")
	e := os.WriteFile(target, []byte("package target\n"), 0644)
	assert.FatalOnError(t, e)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "FormatName", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestExtractRenamesSourceWhenOneRemains(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-last-pair/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "FormatName", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	extracted := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/format_name.go",
	)
	assert.StringContains(t, "func FormatName(", extracted)
	renamed := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/summarize_name.go",
	)
	assert.StringContains(t, "func SummarizeName(", renamed)
	_, e = os.Stat(filepath.Join(d, "pkg/target/combined.go"))
	assert.True(t, os.IsNotExist(e))
}

func TestExtractTypeToFile(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-identity/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "Widget", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := service_tester.ReadFixtureFile(t, d, "pkg/target/widget.go")
	assert.StringContains(t, "type Widget struct", extracted)
	assert.StringContains(t, "// Widget holds a display name.", extracted)
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/combined.go")
	assert.True(t, !strings.Contains(source, "type Widget struct"))
	assert.StringContains(t, "func (w *Widget) Trimmed(", source)
	assert.StringContains(t, "func NewWidget(", source)
}

func TestExtractLeavesTypeWithoutRename(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-identity/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "NewWidget", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/combined.go")
	assert.StringContains(t, "type Widget struct", source)
	assert.StringContains(t, "func (w *Widget) Trimmed(", source)
}

func TestExtractFunctionRenamesSourceToType(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-type-pair/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "NewWidget", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/new_widget.go",
	)
	assert.StringContains(t, "func NewWidget(", extracted)
	renamed := service_tester.ReadFixtureFile(t, d, "pkg/target/widget.go")
	assert.StringContains(t, "type Widget struct", renamed)
	_, e = os.Stat(filepath.Join(d, "pkg/target/combined.go"))
	assert.True(t, os.IsNotExist(e))
}

func TestExtractTypeRenamesSourceToFunction(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-type-pair/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "Widget", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := service_tester.ReadFixtureFile(t, d, "pkg/target/widget.go")
	assert.StringContains(t, "type Widget struct", extracted)
	renamed := service_tester.ReadFixtureFile(t, d, "pkg/target/new_widget.go")
	assert.StringContains(t, "func NewWidget(", renamed)
	_, e = os.Stat(filepath.Join(d, "pkg/target/combined.go"))
	assert.True(t, os.IsNotExist(e))
}

func TestExtractGroupedTypeRefuses(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-type-group/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "Alpha", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "type group")
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/combined.go")
	assert.StringContains(t, "Alpha struct{}", source)
}

func TestExtractTypeRefusesEmptyFile(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-type-single/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/only.go", "Only", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "empty file")
}

func TestExtractLeavesVariableBehind(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-var-companion/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "Register", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := service_tester.ReadFixtureFile(t, d, "pkg/target/register.go")
	assert.StringContains(t, "func Register(", extracted)
	renamed := service_tester.ReadFixtureFile(t, d, "pkg/target/count.go")
	assert.StringContains(t, "var registry", renamed)
	assert.StringContains(t, "func Count(", renamed)
	_, e = os.Stat(filepath.Join(d, "pkg/target/combined.go"))
	assert.True(t, os.IsNotExist(e))
}

func TestExtractRefusesEmptyFile(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-single/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/only.go", "OnlyFunction", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "empty file")
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/only.go")
	assert.StringContains(t, "func OnlyFunction(", source)
}

func TestExtractFromTestFileKeepsItsName(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-test-file/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/target_test.go", "newWidget", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/target_test.go")
	assert.StringContains(t, "func TestWidgetName(", source)
	assert.True(t, !strings.Contains(source, "func newWidget("))
}

func TestExtractFromTestFileCarriesRenamedImport(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-test-file/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/target_test.go", "newWidget", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/new_widget.go",
	)
	assert.StringContains(t, "func newWidget(", extracted)
	assert.StringContains(t, "example/pkg/go-widget", extracted)
}

func TestExtractFromTestFileCarriesBuildConstraint(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-test-file/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/target_test.go", "newWidget", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/new_widget.go",
	)
	assert.Prefix(t, "//go:build ci", extracted)
}

func TestExtractResolvesTagGatedDependency(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-tagged-dependency/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "Describe", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := service_tester.ReadFixtureFile(t, d, "pkg/target/describe.go")
	assert.StringContains(t, "func Describe(", extracted)
	assert.StringContains(t, "example/pkg/gated", extracted)
}
