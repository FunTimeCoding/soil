package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractFunctionNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-function/src"),
	)
	s := service_tester.NewService()
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
	s := service_tester.NewService()
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
	s := service_tester.NewService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "FormatName", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestExtractFunctionRenamesSourceToType(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-type-pair/src"),
	)
	s := service_tester.NewService()
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
	s := service_tester.NewService()
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
	s := service_tester.NewService()
	r, e := s.ExtractToFile(d, "pkg/target/combined.go", "Alfa", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "type group")
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/combined.go")
	assert.StringContains(t, "Alfa  struct{}", source)
}

func TestExtractTypeRefusesEmptyFile(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-type-single/src"),
	)
	s := service_tester.NewService()
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
	s := service_tester.NewService()
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
	s := service_tester.NewService()
	r, e := s.ExtractToFile(d, "pkg/target/only.go", "OnlyFunction", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "empty file")
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/only.go")
	assert.StringContains(t, "func OnlyFunction(", source)
}
