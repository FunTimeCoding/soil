package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestRenameRefusedWhenAFileKeepsChangingBeforeCommit(t *testing.T) {
	d := renameFunctionFixture(t)
	declaration := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/is_generated_header.go",
	)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	s := service_tester.NewService(t)
	s.Commit.BeforeCommit(
		func() {
			testutil.WriteFile(t, d, "pkg/caller/run.go", caller)
		},
	)
	r, e := s.Symbol.Rename(
		t.Context(),
		d,
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
		joinedConcerns(r),
	)
	assert.String(
		t,
		declaration,
		service_tester.ReadFixtureFile(
			t,
			d,
			"pkg/target/is_generated_header.go",
		),
	)
}

func TestRenameRunsAgainAfterAFileChangedOnce(t *testing.T) {
	d := renameFunctionFixture(t)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	s := service_tester.NewService(t)
	edited := false
	s.Commit.BeforeCommit(
		func() {
			if edited {
				return
			}

			edited = true
			testutil.WriteFile(t, d, "pkg/caller/run.go", caller)
		},
	)
	r, e := s.Symbol.Rename(
		t.Context(),
		d,
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
		joinedConcerns(r),
	)
	assert.StringContains(
		t,
		"func IsGenerated(",
		service_tester.ReadFixtureFile(t, d, "pkg/target/is_generated.go"),
	)
}

func TestRenameCommitsPastAChangeOutsideTheReach(t *testing.T) {
	d := renameFunctionFixture(t)
	s := service_tester.NewService(t)
	s.Commit.BeforeCommit(
		func() {
			testutil.WriteFile(
				t,
				d,
				"pkg/other/other.go",
				"package other\n\nfunc Other() {}\n",
			)
		},
	)
	r, e := s.Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringNotContains(t, "pkg/other/other.go", joinedConcerns(r))
	assert.StringContains(
		t,
		"func IsGenerated(",
		service_tester.ReadFixtureFile(t, d, "pkg/target/is_generated.go"),
	)
}

func TestRenameRefusedWhenAPackageStartsImportingTheTarget(t *testing.T) {
	d := renameFunctionFixture(t)
	s := service_tester.NewService(t)
	s.Commit.BeforeCommit(
		func() {
			testutil.WriteFile(
				t,
				d,
				"pkg/other/other.go",
				"package other\n\nimport \"example/pkg/target\"\n\nfunc Other() bool {\n\treturn target.IsGeneratedHeader(\"\")\n}\n",
			)
		},
	)
	r, e := s.Symbol.Rename(
		t.Context(),
		d,
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
		joinedConcerns(r),
	)
}

func TestRenameRefusedWhenAReplacingModuleChangesBeforeCommit(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	source := service_tester.ReadFixtureFile(t, library, "lib.go")
	consumer := service_tester.ReadFixtureFile(t, user, "pkg/user/user.go")
	s := replacingService(t, library, user)
	s.Commit.BeforeCommit(
		func() {
			testutil.WriteFile(t, user, "pkg/user/user.go", consumer)
		},
	)
	r, e := s.Symbol.Rename(
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
		filepath.Join(user, "pkg/user/user.go: changed since it was read"),
		joinedConcerns(r),
	)
	assert.String(
		t,
		source,
		service_tester.ReadFixtureFile(t, library, "lib.go"),
	)
	assert.StringContains(
		t,
		"lib.Log(\"a\")",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
}

func TestRenameCommitsPastAReplacingModuleChangeOutsideTheReach(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	s := replacingService(t, library, user)
	s.Commit.BeforeCommit(
		func() {
			testutil.WriteFile(
				t,
				user,
				"pkg/alfa/alfa.go",
				"package alfa\n\nfunc Alfa() string {\n\treturn \"changed\"\n}\n",
			)
		},
	)
	r, e := s.Symbol.Rename(
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
	assert.StringContains(
		t,
		"lib.Write(\"a\")",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
}

func TestRenameCommitsWhenNothingChanged(t *testing.T) {
	d := renameFunctionFixture(t)
	called := false
	s := service_tester.NewService(t)
	s.Commit.BeforeCommit(
		func() {
			called = true
		},
	)
	r, e := s.Symbol.Rename(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "pkg/target/is_generated.go"),
	)
}
