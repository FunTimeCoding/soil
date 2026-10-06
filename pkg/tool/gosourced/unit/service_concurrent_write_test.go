package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestRenameRefusedWhenAFileChangesBeforeCommit(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-function/src"),
	)
	declaration := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/is_generated_header.go",
	)
	s := testService()
	s.BeforeCommit(
		func() {
			testutil.WriteFile(
				t,
				d,
				"pkg/other/other.go",
				"package other\n\nfunc Other() {}\n",
			)
		},
	)
	r, e := s.Rename(
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
		"pkg/other/other.go: changed since it was read - nothing written, run again",
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

func TestRenameRefusedWhenAReplacingModuleChangesBeforeCommit(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	source := service_tester.ReadFixtureFile(t, library, "lib.go")
	s := replacingService(library, user)
	s.BeforeCommit(
		func() {
			testutil.WriteFile(
				t,
				user,
				"pkg/alfa/alfa.go",
				"package alfa\n\nfunc Alfa() string {\n\treturn \"changed\"\n}\n",
			)
		},
	)
	r, e := s.Rename(library, "other.test/lib", "Log", "Write", "", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.StringContains(
		t,
		filepath.Join(user, "pkg/alfa/alfa.go: changed since it was read"),
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

func TestRenameCommitsWhenNothingChanged(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-function/src"),
	)
	called := false
	s := testService()
	s.BeforeCommit(
		func() {
			called = true
		},
	)
	r, e := s.Rename(
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
		service_tester.ReadFixtureFile(
			t,
			d,
			"pkg/target/is_generated_header.go",
		),
	)
}
