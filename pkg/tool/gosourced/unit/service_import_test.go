package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"strings"
	"testing"
)

func TestAddImportToGrouped(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("import-grouped/src"))
	r, e := o.File.AddImport(
		t.Context(),
		o.Directory,
		"pkg/target/example.go",
		"os",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/target/example.go")
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
	o := service_tester.New(t, service_tester.WithFixture("import-empty/src"))
	r, e := o.File.AddImport(
		t.Context(),
		o.Directory,
		"pkg/target/example.go",
		"fmt",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/target/example.go")
	assert.StringContains(t, "\"fmt\"", source)
}

func TestAddImportWithAlias(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("import-grouped/src"))
	r, e := o.File.AddImport(
		t.Context(),
		o.Directory,
		"pkg/target/example.go",
		"path/filepath",
		"fp",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/target/example.go")
	assert.StringContains(t, "fp \"path/filepath\"", source)
}

func TestRemoveImport(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("import-grouped/src"))
	r, e := o.File.RemoveImport(
		t.Context(),
		o.Directory,
		"pkg/target/example.go",
		"strings",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/target/example.go")
	assert.True(t, !strings.Contains(source, "\"strings\""))
	assert.StringContains(t, "strings.TrimSpace", source)
	assert.StringContains(t, "\"fmt\"", source)
}

func TestRemoveImportNotFound(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("import-grouped/src"))
	r, e := o.File.RemoveImport(
		t.Context(),
		o.Directory,
		"pkg/target/example.go",
		"os",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestAddImportRefusesFileOutsideModule(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("import-grouped/src"))
	other := service_tester.New(
		t,
		service_tester.WithFixture("import-grouped/src"),
	)
	source := other.Read("pkg/target/example.go")
	r, e := o.File.AddImport(
		t.Context(),
		o.Directory,
		filepath.Join(other.Directory, "pkg/target/example.go"),
		"os",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "outside the active module")
	assert.String(t, source, other.Read("pkg/target/example.go"))
}
