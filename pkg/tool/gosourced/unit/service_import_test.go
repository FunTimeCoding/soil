package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"strings"
	"testing"
)

func TestAddImportToGrouped(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("import-grouped/src"))
	s := testService()
	r, e := s.AddImport(d, "pkg/target/example.go", "os", "", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := readFixtureFile(t, d, "pkg/target/example.go")
	assert.StringContains(t, "\"os\"", source)
	assert.StringContains(t, "\"fmt\"", source)
	assert.StringContains(t, "// Example formats a trimmed value.", source)
	assert.StringContains(t, "} // example trailing", source)
	assert.True(
		t,
		strings.Index(source, "// example trailing") >
			strings.Index(source, "\"os\""),
	)
}

func TestAddImportToEmpty(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("import-empty/src"))
	s := testService()
	r, e := s.AddImport(d, "pkg/target/example.go", "fmt", "", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := readFixtureFile(t, d, "pkg/target/example.go")
	assert.StringContains(t, "\"fmt\"", source)
}

func TestAddImportWithAlias(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("import-grouped/src"))
	s := testService()
	r, e := s.AddImport(
		d,
		"pkg/target/example.go",
		"path/filepath",
		"fp",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := readFixtureFile(t, d, "pkg/target/example.go")
	assert.StringContains(t, "fp \"path/filepath\"", source)
}

func TestRemoveImport(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("import-grouped/src"))
	s := testService()
	r, e := s.RemoveImport(d, "pkg/target/example.go", "strings", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := readFixtureFile(t, d, "pkg/target/example.go")
	assert.True(t, !strings.Contains(source, "\"strings\""))
	assert.StringContains(t, "strings.TrimSpace", source)
	assert.StringContains(t, "\"fmt\"", source)
}

func TestRemoveImportNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("import-grouped/src"))
	s := testService()
	r, e := s.RemoveImport(d, "pkg/target/example.go", "os", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}
