package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"path/filepath"
	"testing"
)

func TestRenameRewritesReplacingModule(t *testing.T) {
	library, o := module_tester.New(t).ReplacedLibrary("rename-replacer")
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
	source := system_tester.New(t).Read(library, "lib.go")
	assert.StringContains(t, "func Write(prefix string) string {", source)
	assert.StringContains(t, "return Write(\"other\")", source)
	assert.StringContains(
		t,
		"return lib.Write(\"a\")",
		o.Read("pkg/user/user.go"),
	)
	assert.StringContains(
		t,
		filepath.Join(o.Directory, "pkg/user/user.go: Log → Write"),
		rewrite_tester.JoinedConcerns(r),
	)
}

func TestRenameFieldReachesReplacingModuleThroughAnotherPackage(t *testing.T) {
	library, o := module_tester.New(t).ReplacedLibrary("rename-replacer")
	r, e := o.Symbol.Rename(
		t.Context(),
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
		o.Read("pkg/user/user.go"),
	)
}

func TestRenameDryRunLeavesReplacingModuleUntouched(t *testing.T) {
	library, o := module_tester.New(t).ReplacedLibrary("rename-replacer")
	source := o.Read("pkg/user/user.go")
	r, e := o.Symbol.Rename(
		t.Context(),
		library,
		"other.test/lib",
		"Log",
		"Write",
		"",
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, source, o.Read("pkg/user/user.go"))
	assert.Integer(t, 3, len(r.Entries))
	assert.String(
		t,
		filepath.Join(o.Directory, "pkg/user/user.go"),
		r.Entries[2].Path,
	)
	assert.True(t, r.Entries[2].Planned)
}

func TestRenameRefusesUnexportingWhatReplacingModuleUses(t *testing.T) {
	library, o := module_tester.New(t).ReplacedLibrary("rename-replacer")
	source := system_tester.New(t).Read(library, "lib.go")
	r, e := o.Symbol.Rename(
		t.Context(),
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
			o.Directory,
			"pkg/user/user.go: example/pkg/user.Log would lose access",
		),
		rewrite_tester.JoinedConcerns(r),
	)
	assert.String(t, source, system_tester.New(t).Read(library, "lib.go"))
}

func TestRenameRefusedWhenReplacingModuleDoesNotTypeCheck(t *testing.T) {
	library, o := module_tester.New(t).ReplacedLibrary("rename-replacer")
	system_tester.New(t).Write(
		o.Directory,
		"pkg/broken/broken.go",
		"package broken\n\nimport \"other.test/lib\"\n\nfunc Broken() string {\n\treturn lib.Log(missing)\n}\n",
	)
	source := system_tester.New(t).Read(library, "lib.go")
	_, e := o.Symbol.Rename(
		t.Context(),
		library,
		"other.test/lib",
		"Log",
		"Write",
		"",
		false,
	)
	assert.Error(t, e)
	assert.String(t, source, system_tester.New(t).Read(library, "lib.go"))
}

func TestRenameSkipsReplacingModuleThatDoesNotReach(t *testing.T) {
	library, o := module_tester.New(t).ReplacedLibrary("rename-replacer")
	r, e := o.Symbol.Rename(
		t.Context(),
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
		system_tester.New(t).Read(library, "unused/unused.go"),
	)
}
