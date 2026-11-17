package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestRenameRefusedWhenFileKeepsChangingBeforeCommit(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	declaration := o.Read("pkg/target/is_generated_header.go")
	caller := o.Read("pkg/caller/run.go")
	o.Commit.BeforeCommit(
		func() {
			system_tester.Write(t, o.Directory, "pkg/caller/run.go", caller)
		},
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.String(
		t,
		"pkg/caller/run.go: changed since it was read - nothing written, run again\npkg/caller/run.go: changed while the call ran - ran again\npkg/caller/run.go: changed while the call ran - ran again",
		rewrite_tester.JoinedConcerns(r),
	)
	assert.String(t, declaration, o.Read("pkg/target/is_generated_header.go"))
}

func TestRenameRunsAgainAfterFileChangedOnce(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	caller := o.Read("pkg/caller/run.go")
	edited := false
	o.Commit.BeforeCommit(
		func() {
			if edited {
				return
			}

			edited = true
			system_tester.Write(t, o.Directory, "pkg/caller/run.go", caller)
		},
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"pkg/caller/run.go: changed while the call ran - ran again",
		rewrite_tester.JoinedConcerns(r),
	)
	assert.StringContains(
		t,
		"func IsGenerated(",
		o.Read("pkg/target/is_generated.go"),
	)
}

func TestRenameCommitsPastChangeOutsideReach(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	o.Commit.BeforeCommit(
		func() {
			system_tester.Write(
				t,
				o.Directory,
				"pkg/other/other.go",
				"package other\n\nfunc Other() {}\n",
			)
		},
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringNotContains(
		t,
		"pkg/other/other.go",
		rewrite_tester.JoinedConcerns(r),
	)
	assert.StringContains(
		t,
		"func IsGenerated(",
		o.Read("pkg/target/is_generated.go"),
	)
}

func TestRenameRefusedWhenPackageStartsImportingTarget(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	o.Commit.BeforeCommit(
		func() {
			system_tester.Write(
				t,
				o.Directory,
				"pkg/other/other.go",
				"package other\n\nimport \"example/pkg/target\"\n\nfunc Other() bool {\n\treturn target.IsGeneratedHeader(\"\")\n}\n",
			)
		},
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.StringContains(
		t,
		"pkg/other/other.go: changed since it was read",
		rewrite_tester.JoinedConcerns(r),
	)
}

func TestRenameRefusedWhenReplacingModuleChangesBeforeCommit(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "rename-replacer")
	source := system_tester.Read(t, library, "lib.go")
	consumer := o.Read("pkg/user/user.go")
	o.Commit.BeforeCommit(
		func() {
			system_tester.Write(t, o.Directory, "pkg/user/user.go", consumer)
		},
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		library,
		"other.test/lib",
		"Log",
		"Write",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.StringContains(
		t,
		filepath.Join(
			o.Directory,
			"pkg/user/user.go: changed since it was read",
		),
		rewrite_tester.JoinedConcerns(r),
	)
	assert.String(t, source, system_tester.Read(t, library, "lib.go"))
	assert.StringContains(t, "lib.Log(\"a\")", o.Read("pkg/user/user.go"))
}

func TestRenameCommitsPastReplacingModuleChangeOutsideReach(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "rename-replacer")
	o.Commit.BeforeCommit(
		func() {
			system_tester.Write(
				t,
				o.Directory,
				"pkg/alfa/alfa.go",
				"package alfa\n\nfunc Alfa() string {\n\treturn \"changed\"\n}\n",
			)
		},
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		library,
		"other.test/lib",
		"Log",
		"Write",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(t, "lib.Write(\"a\")", o.Read("pkg/user/user.go"))
}

func TestRenameCommitsWhenNothingChanged(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	called := false
	o.Commit.BeforeCommit(
		func() {
			called = true
		},
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, called)
	assert.StringContains(
		t,
		"func IsGenerated(",
		o.Read("pkg/target/is_generated.go"),
	)
}
