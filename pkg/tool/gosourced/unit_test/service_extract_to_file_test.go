package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtractFunction(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("extract-function/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "FormatName")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := readFixtureFile(t, d, "pkg/target/format_name.go")
	assert.StringContains(t, "func FormatName(", extracted)
	assert.StringContains(t, "fmt.Sprintf", extracted)
	source := readFixtureFile(t, d, "pkg/target/combined.go")
	assert.True(t, !strings.Contains(source, "func FormatName("))
	assert.StringContains(t, "func TrimName(", source)
	assert.StringContains(t, "func PlainName(", source)
}

func TestExtractFunctionRemovesUnusedImport(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("extract-function/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "TrimName")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := readFixtureFile(t, d, "pkg/target/trim_name.go")
	assert.StringContains(t, "func TrimName(", extracted)
	assert.StringContains(t, "strings", extracted)
	source := readFixtureFile(t, d, "pkg/target/combined.go")
	assert.True(t, !strings.Contains(source, "strings"))
	assert.StringContains(t, "fmt", source)
}

func TestExtractFunctionNoImports(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("extract-function/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "PlainName")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := readFixtureFile(t, d, "pkg/target/plain_name.go")
	assert.StringContains(t, "func PlainName(", extracted)
	assert.True(t, !strings.Contains(extracted, "import"))
}

func TestExtractFunctionNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("extract-function/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "Missing")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestExtractFunctionFileNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("extract-function/src"),
	)
	s := testService()
	_, e := s.ExtractToFile(d, "pkg/target/missing.go", "Something")
	assert.True(t, e != nil)
}

func TestExtractFunctionTargetExists(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("extract-function/src"),
	)
	target := filepath.Join(d, "pkg/target/format_name.go")
	e := os.WriteFile(target, []byte("package target\n"), 0644)
	assert.FatalOnError(t, e)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "FormatName")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestExtractRenamesSourceWhenOneRemains(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("extract-last-pair/src"),
	)
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "FormatName")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	extracted := readFixtureFile(t, d, "pkg/target/format_name.go")
	assert.StringContains(t, "func FormatName(", extracted)
	renamed := readFixtureFile(t, d, "pkg/target/summarize_name.go")
	assert.StringContains(t, "func SummarizeName(", renamed)
	_, e = os.Stat(filepath.Join(d, "pkg/target/combined.go"))
	assert.True(t, os.IsNotExist(e))
}

func TestExtractRefusesEmptyFile(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("extract-single/src"))
	s := testService()
	r, e := s.ExtractToFile(d, "pkg/target/only.go", "OnlyFunction")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "empty file")
	source := readFixtureFile(t, d, "pkg/target/only.go")
	assert.StringContains(t, "func OnlyFunction(", source)
}
