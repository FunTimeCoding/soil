package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestMovePackageRetargetsAReplacingModule(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "package-replacer")
	r, e := replacingService(library, user).MovePackage(
		library,
		"other.test/lib/shape",
		"other.test/lib/geometry/shape",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(
		t,
		system.FileExists(filepath.Join(library, "geometry/shape/shape.go")),
	)
	assert.String(
		t,
		"package user\n\nimport \"other.test/lib/geometry/shape\"\n\nfunc Width() int {\n\treturn shape.New().Width\n}\n",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
	assert.StringContains(
		t,
		"\"other.test/lib/geometry/shape\"",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user_test.go"),
	)
	assert.StringContains(
		t,
		"import \"other.test/lib/geometry/shape/sub\"",
		service_tester.ReadFixtureFile(t, user, "pkg/deep/deep.go"),
	)
	assert.StringContains(
		t,
		"import figure \"other.test/lib/geometry/shape\"",
		service_tester.ReadFixtureFile(t, user, "pkg/aliased/aliased.go"),
	)
}

func TestRenamePackageRenamesQualifiersInAReplacingModule(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "package-replacer")
	r, e := replacingService(library, user).RenamePackage(
		library,
		"other.test/lib/shape",
		"form",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package user\n\nimport \"other.test/lib/form\"\n\nfunc Width() int {\n\treturn form.New().Width\n}\n",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
	assert.StringContains(
		t,
		"if Width() != form.New().Width {",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user_test.go"),
	)
	assert.StringContains(
		t,
		"import \"other.test/lib/form/sub\"",
		service_tester.ReadFixtureFile(t, user, "pkg/deep/deep.go"),
	)
	aliased := service_tester.ReadFixtureFile(t, user, "pkg/aliased/aliased.go")
	assert.StringContains(t, "import figure \"other.test/lib/form\"", aliased)
	assert.StringContains(t, "return figure.New()", aliased)
}

func TestRenamePackageClauseRenamesQualifiersInAReplacingModule(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "package-replacer")
	r, e := replacingService(library, user).RenamePackageClause(
		library,
		"other.test/lib/shape",
		"form",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(t, user, "pkg/user/user.go")
	assertFormatted(t, source)
	assert.StringContains(t, "\"other.test/lib/shape\"", source)
	assert.StringContains(t, "return form.New().Width", source)
	assert.StringContains(
		t,
		"return figure.New()",
		service_tester.ReadFixtureFile(t, user, "pkg/aliased/aliased.go"),
	)
}

func TestRenamePackageRefusedWhenTheNameIsTakenInAReplacingModule(
	t *testing.T,
) {
	library, user := prepareReplacedLibrary(t, "package-replacer")
	testutil.WriteFile(
		t,
		user,
		"pkg/clash/clash.go",
		"package clash\n\nimport \"other.test/lib/shape\"\n\nfunc form() int {\n\treturn shape.New().Width\n}\n",
	)
	r, e := replacingService(library, user).RenamePackage(
		library,
		"other.test/lib/shape",
		"form",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "form is taken")
	assert.True(t, system.FileExists(filepath.Join(library, "shape/shape.go")))
	assert.StringContains(
		t,
		"import \"other.test/lib/shape\"",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
}

func TestMovePackageReachesAReplacingModuleThroughASubpackage(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "package-replacer")
	assert.FatalOnError(t, os.RemoveAll(filepath.Join(user, "pkg/user")))
	assert.FatalOnError(t, os.RemoveAll(filepath.Join(user, "pkg/aliased")))
	r, e := replacingService(library, user).MovePackage(
		library,
		"other.test/lib/shape",
		"other.test/lib/geometry/shape",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"import \"other.test/lib/geometry/shape/sub\"",
		service_tester.ReadFixtureFile(t, user, "pkg/deep/deep.go"),
	)
}

func TestMovePackageDryRunLeavesAReplacingModuleUntouched(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "package-replacer")
	before := treeOf(t, user)
	r, e := replacingService(library, user).MovePackage(
		library,
		"other.test/lib/shape",
		"other.test/lib/geometry/shape",
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, before, treeOf(t, user))
	assert.True(t, system.FileExists(filepath.Join(library, "shape/shape.go")))
	assert.StringContains(
		t,
		filepath.Join(user, "pkg/deep/deep.go"),
		joinedConcerns(r),
	)
}
