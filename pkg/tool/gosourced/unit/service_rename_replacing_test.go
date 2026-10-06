package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestRenameRewritesAReplacingModule(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	r, e := replacingService(library, user).Rename(
		library,
		"other.test/lib",
		"Log",
		"Write",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(t, library, "lib.go")
	assert.StringContains(t, "func Write(prefix string) string {", source)
	assert.StringContains(t, "return Write(\"other\")", source)
	assert.StringContains(
		t,
		"return lib.Write(\"a\")",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
	assert.StringContains(
		t,
		filepath.Join(user, "pkg/user/user.go: Log → Write"),
		joinedConcerns(r),
	)
}

func TestRenameFieldReachesAReplacingModuleThroughAnotherPackage(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	r, e := replacingService(library, user).Rename(
		library,
		"other.test/lib/shape",
		"Width",
		"Size",
		"Shape",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"return maker.Make().Size",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
}

func TestRenameDryRunLeavesAReplacingModuleUntouched(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	source := service_tester.ReadFixtureFile(t, user, "pkg/user/user.go")
	r, e := replacingService(library, user).Rename(
		library,
		"other.test/lib",
		"Log",
		"Write",
		"",
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		source,
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
	assert.Integer(t, 3, len(r.Entries))
	assert.String(t, filepath.Join(user, "pkg/user/user.go"), r.Entries[2].Path)
	assert.True(t, r.Entries[2].Planned)
}

func TestRenameRefusesUnexportingWhatAReplacingModuleUses(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	source := service_tester.ReadFixtureFile(t, library, "lib.go")
	r, e := replacingService(library, user).Rename(
		library,
		"other.test/lib",
		"Log",
		"log",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.StringContains(
		t,
		filepath.Join(
			user,
			"pkg/user/user.go: example/pkg/user.Log would lose access",
		),
		joinedConcerns(r),
	)
	assert.String(
		t,
		source,
		service_tester.ReadFixtureFile(t, library, "lib.go"),
	)
}

func TestRenameRefusedWhenAReplacingModuleDoesNotTypeCheck(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	testutil.WriteFile(
		t,
		user,
		"pkg/broken/broken.go",
		"package broken\n\nimport \"other.test/lib\"\n\nfunc Broken() string {\n\treturn lib.Log(missing)\n}\n",
	)
	source := service_tester.ReadFixtureFile(t, library, "lib.go")
	_, e := replacingService(library, user).Rename(
		library,
		"other.test/lib",
		"Log",
		"Write",
		"",
		false,
	)
	assert.Error(t, e)
	assert.String(
		t,
		source,
		service_tester.ReadFixtureFile(t, library, "lib.go"),
	)
}

func TestRenameSkipsAReplacingModuleThatDoesNotReach(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	r, e := replacingService(library, user).Rename(
		library,
		"other.test/lib/unused",
		"Spare",
		"Extra",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"func Extra() string {",
		service_tester.ReadFixtureFile(t, library, "unused/unused.go"),
	)
}
