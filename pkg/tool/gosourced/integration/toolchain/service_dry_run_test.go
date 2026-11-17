package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestDryRunLeavesTreeUntouched(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("batch-file/src"))
	source := o.Read("pkg/target/constant.go")
	run := o.Read("pkg/target/run.go")
	target := o.Read("pkg/target/constant/constant.go")
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, source, o.Read("pkg/target/constant.go"))
	assert.String(t, run, o.Read("pkg/target/run.go"))
	assert.String(t, target, o.Read("pkg/target/constant/constant.go"))
}

func TestDryRunReportsWhatRealRunReports(t *testing.T) {
	planned := service_tester.New(
		t,
		service_tester.WithFixture("batch-file/src"),
	)
	dry, e := planned.Symbol.MoveSymbols(
		t.Context(),
		planned.Directory,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		true,
	)
	assert.FatalOnError(t, e)
	applied := service_tester.New(
		t,
		service_tester.WithFixture("batch-file/src"),
	)
	real, f := applied.Symbol.MoveSymbols(
		t.Context(),
		applied.Directory,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, f)
	assert.Strings(
		t,
		service_tester.ConcernText(real),
		service_tester.ConcernText(dry),
	)
}

func TestDryRunMarksConcernsPlanned(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("batch-file/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		true,
	)
	assert.FatalOnError(t, e)
	assert.Greater(t, 0, len(r.Entries))

	for _, c := range r.Entries {
		assert.True(t, c.Planned)
		assert.False(t, c.Fixed)
	}
}

func TestDryRunRenameLeavesFilesUntouched(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("batch-enum/src"))
	before := o.Read("pkg/target/run.go")
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Run",
		"Execute",
		"",
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, before, o.Read("pkg/target/run.go"))
	assert.True(t, strings.Contains(before, "func Run()"))
}
