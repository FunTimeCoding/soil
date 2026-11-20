package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/fix_tester"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"testing"
)

func TestFixRefusedWhenReachedPackageChangesBeforeCommit(t *testing.T) {
	m := module_tester.New(t)
	directory := m.NamesakeExported()
	o := m.Options(directory, false, false, "./alfa/...")
	o.BeforeCommit = func() {
		system_tester.New(t).Write(
			directory,
			"bravo/extra.go",
			"package bravo\n",
		)
	}
	r := output.NewResultsWithDirectory(directory)
	m.RunDefault(o, r)
	assert.String(
		t,
		"package alfa\n\nfunc SendMsg() string {\n\treturn \"x\"\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "alfa", "send.go")),
	)
	assert.Integer(t, 0, len(fix_tester.Applied(r.Entries)))
	last := r.Entries[len(r.Entries)-1]
	assert.String(t, "bravo/extra.go", last.Path)
	assert.String(
		t,
		"changed since it was read - nothing written, run again",
		last.Text,
	)
}

func TestFixCommitsPastChangeOutsideReach(t *testing.T) {
	m := module_tester.New(t)
	directory := m.NamesakeExported()
	o := m.Options(directory, false, false, "./alfa/...")
	o.BeforeCommit = func() {
		system_tester.New(t).Write(
			directory,
			"delta/other.go",
			"package delta\n\nfunc Other() string {\n\treturn \"changed\"\n}\n",
		)
	}
	m.RunDefault(o, output.NewResultsWithDirectory(directory))
	assert.String(
		t,
		"package alfa\n\nfunc SendMessage() string {\n\treturn \"x\"\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "alfa", "send.go")),
	)
}

func TestFixRefusedWhenReplacingModuleChangesBeforeCommit(t *testing.T) {
	m := module_tester.New(t)
	library := m.ReplacedLibrary()
	user := m.ReplacingUser(
		library,
		"package user\n\nimport \"other.test/lib/alfa\"\n\nfunc Use() string {\n\treturn alfa.SendMsg()\n}\n",
	)
	o := m.Options(library, false, false, "./alfa/...")
	o.Replacing = []string{user}
	o.BeforeCommit = func() {
		system_tester.New(t).Write(user, "user/extra.go", "package user\n")
	}
	m.RunDefault(o, output.NewResultsWithDirectory(library))
	assert.String(
		t,
		"package user\n\nimport \"other.test/lib/alfa\"\n\nfunc Use() string {\n\treturn alfa.SendMsg()\n}\n",
		testutil.ReadFile(t, filepath.Join(user, "user", "use.go")),
	)
	assert.String(
		t,
		"package alfa\n\nfunc SendMsg() string {\n\treturn \"x\"\n}\n",
		testutil.ReadFile(t, filepath.Join(library, "alfa", "send.go")),
	)
}
