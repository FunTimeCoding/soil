package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestMovePackageRetargetsReplacingModule(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "package-replacer")
	r, e := o.Packages.MovePackage(
		t.Context(),
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
		o.Read("pkg/user/user.go"),
	)
	assert.StringContains(
		t,
		"\"other.test/lib/geometry/shape\"",
		o.Read("pkg/user/user_test.go"),
	)
	assert.StringContains(
		t,
		"import \"other.test/lib/geometry/shape/sub\"",
		o.Read("pkg/deep/deep.go"),
	)
	assert.StringContains(
		t,
		"import figure \"other.test/lib/geometry/shape\"",
		o.Read("pkg/aliased/aliased.go"),
	)
}

func TestRenamePackageRenamesQualifiersInReplacingModule(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "package-replacer")
	r, e := o.Packages.RenamePackage(
		t.Context(),
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
		o.Read("pkg/user/user.go"),
	)
	assert.StringContains(
		t,
		"if Width() != form.New().Width {",
		o.Read("pkg/user/user_test.go"),
	)
	assert.StringContains(
		t,
		"import \"other.test/lib/form/sub\"",
		o.Read("pkg/deep/deep.go"),
	)
	aliased := o.Read("pkg/aliased/aliased.go")
	assert.StringContains(t, "import figure \"other.test/lib/form\"", aliased)
	assert.StringContains(t, "return figure.New()", aliased)
}

func TestRenamePackageClauseRenamesQualifiersInReplacingModule(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "package-replacer")
	r, e := o.Packages.RenamePackageClause(
		t.Context(),
		library,
		"other.test/lib/shape",
		"form",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/user/user.go")
	rewrite_tester.Formatted(t, source)
	assert.StringContains(t, "\"other.test/lib/shape\"", source)
	assert.StringContains(t, "return form.New().Width", source)
	assert.StringContains(
		t,
		"return figure.New()",
		o.Read("pkg/aliased/aliased.go"),
	)
}

func TestRenamePackageRefusedWhenNameIsTakenInReplacingModule(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "package-replacer")
	system_tester.Write(
		t,
		o.Directory,
		"pkg/clash/clash.go",
		"package clash\n\nimport \"other.test/lib/shape\"\n\nfunc form() int {\n\treturn shape.New().Width\n}\n",
	)
	r, e := o.Packages.RenamePackage(
		t.Context(),
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
		o.Read("pkg/user/user.go"),
	)
}

func TestMovePackageReachesReplacingModuleThroughSubpackage(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "package-replacer")
	assert.FatalOnError(t, os.RemoveAll(filepath.Join(o.Directory, "pkg/user")))
	assert.FatalOnError(
		t,
		os.RemoveAll(filepath.Join(o.Directory, "pkg/aliased")),
	)
	r, e := o.Packages.MovePackage(
		t.Context(),
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
		o.Read("pkg/deep/deep.go"),
	)
}

func TestMovePackageDryRunLeavesReplacingModuleUntouched(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "package-replacer")
	before := rewrite_tester.DirectoryTree(t, o.Directory)
	r, e := o.Packages.MovePackage(
		t.Context(),
		library,
		"other.test/lib/shape",
		"other.test/lib/geometry/shape",
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, before, rewrite_tester.DirectoryTree(t, o.Directory))
	assert.True(t, system.FileExists(filepath.Join(library, "shape/shape.go")))
	assert.StringContains(
		t,
		filepath.Join(o.Directory, "pkg/deep/deep.go"),
		rewrite_tester.JoinedConcerns(r),
	)
}
