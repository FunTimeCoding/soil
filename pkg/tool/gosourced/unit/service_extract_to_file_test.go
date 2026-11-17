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
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-function/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"Missing",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestExtractFunctionFileNotFound(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-function/src"),
	)
	_, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/missing.go",
		"Something",
		false,
	)
	assert.True(t, e != nil)
}

func TestExtractFunctionTargetExists(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-function/src"),
	)
	target := filepath.Join(o.Directory, "pkg/target/format_name.go")
	e := os.WriteFile(target, []byte("package target\n"), 0644)
	assert.FatalOnError(t, e)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"FormatName",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestExtractFunctionRenamesSourceToType(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-type-pair/src"),
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
	extracted := o.Read("pkg/target/new_widget.go")
	assert.StringContains(t, "func NewWidget(", extracted)
	renamed := o.Read("pkg/target/widget.go")
	assert.StringContains(t, "type Widget struct", renamed)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/target/combined.go"))
	assert.True(t, os.IsNotExist(e))
}

func TestExtractTypeRenamesSourceToFunction(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-type-pair/src"),
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
	renamed := o.Read("pkg/target/new_widget.go")
	assert.StringContains(t, "func NewWidget(", renamed)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/target/combined.go"))
	assert.True(t, os.IsNotExist(e))
}

func TestExtractGroupedTypeRefuses(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-type-group/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"Alfa",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "type group")
	source := o.Read("pkg/target/combined.go")
	assert.StringContains(t, "Alfa  struct{}", source)
}

func TestExtractTypeRefusesEmptyFile(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-type-single/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/only.go",
		"Only",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "empty file")
}

func TestExtractLeavesVariableBehind(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-var-companion/src"),
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"Register",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	extracted := o.Read("pkg/target/register.go")
	assert.StringContains(t, "func Register(", extracted)
	renamed := o.Read("pkg/target/count.go")
	assert.StringContains(t, "var registry", renamed)
	assert.StringContains(t, "func Count(", renamed)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/target/combined.go"))
	assert.True(t, os.IsNotExist(e))
}

func TestExtractRefusesEmptyFile(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("extract-single/src"))
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/only.go",
		"OnlyFunction",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "empty file")
	source := o.Read("pkg/target/only.go")
	assert.StringContains(t, "func OnlyFunction(", source)
}
