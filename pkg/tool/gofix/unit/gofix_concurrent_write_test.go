package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/unit/module_tester"
	"path/filepath"
	"testing"
)

func TestFixRefusedWhenAFileChangesBeforeCommit(t *testing.T) {
	directory := module_tester.NamesakeExported(t)
	o := module_tester.Options(t, directory, false, false, "./alfa/...")
	o.BeforeCommit = func() {
		testutil.WriteFile(
			t,
			directory,
			"delta/other.go",
			"package delta\n\nfunc Other() string {\n\treturn \"changed\"\n}\n",
		)
	}
	r := output.NewResultsWithDirectory(directory)
	gofix.RunDefault(o, r)
	assert.String(
		t,
		"package alfa\n\nfunc SendMsg() string {\n\treturn \"x\"\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "alfa", "send.go")),
	)
	assert.Integer(t, 0, len(filterApplied(r.Entries)))
	last := r.Entries[len(r.Entries)-1]
	assert.String(t, "delta/other.go", last.Path)
	assert.String(
		t,
		"changed since it was read - nothing written, run again",
		last.Text,
	)
}

func TestFixRefusedWhenAReplacingModuleChangesBeforeCommit(t *testing.T) {
	library := module_tester.ReplacedLibrary(t)
	user := module_tester.ReplacingUser(
		t,
		library,
		"package user\n\nimport \"other.test/lib/alfa\"\n\nfunc Use() string {\n\treturn alfa.SendMsg()\n}\n",
	)
	o := module_tester.Options(t, library, false, false, "./alfa/...")
	o.Replacing = []string{user}
	o.BeforeCommit = func() {
		testutil.WriteFile(t, user, "user/extra.go", "package user\n")
	}
	gofix.RunDefault(o, output.NewResultsWithDirectory(library))
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
