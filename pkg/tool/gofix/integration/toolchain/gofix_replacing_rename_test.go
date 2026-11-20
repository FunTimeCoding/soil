package toolchain

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/fix_tester"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"testing"
)

func TestExportedRenameReachesReplacingModule(t *testing.T) {
	m := module_tester.New(t)

	for _, full := range []bool{false, true} {
		library := m.ReplacedLibrary()
		user := m.ReplacingUser(
			library,
			"package user\n\nimport \"other.test/lib/alfa\"\n\nfunc Use() string {\n\treturn alfa.SendMsg()\n}\n",
		)
		o := m.Options(library, false, full, "./alfa/...")
		o.Replacing = []string{user}
		m.RunDefault(o, output.NewResultsWithDirectory(library))
		assert.String(
			t,
			"package user\n\nimport \"other.test/lib/alfa\"\n\nfunc Use() string {\n\treturn alfa.SendMessage()\n}\n",
			testutil.ReadFile(t, filepath.Join(user, "user", "use.go")),
		)
		assert.String(
			t,
			"package alfa\n\nfunc SendMessage() string {\n\treturn \"x\"\n}\n",
			testutil.ReadFile(t, filepath.Join(library, "alfa", "send.go")),
		)
	}
}

func TestExportedRenameDiffLeavesReplacingModule(t *testing.T) {
	m := module_tester.New(t)
	library := m.ReplacedLibrary()
	source := "package user\n\nimport \"other.test/lib/alfa\"\n\nfunc Use() string {\n\treturn alfa.SendMsg()\n}\n"
	user := m.ReplacingUser(library, source)
	o := m.Options(library, true, false, "./alfa/...")
	o.Replacing = []string{user}
	printed := assert.Capture(
		t,
		func() {
			m.RunDefault(o, output.NewResultsWithDirectory(library))
		},
	)
	assert.String(
		t,
		source,
		testutil.ReadFile(t, filepath.Join(user, "user", "use.go")),
	)
	assert.StringContains(
		t,
		fmt.Sprintf(
			"+++ %s\n@@\n \n func Use() string {\n-\treturn alfa.SendMsg()\n+\treturn alfa.SendMessage()\n }\n",
			filepath.Join(user, "user", "use.go"),
		),
		printed,
	)
}

func TestExportedRenameRefusedWhenReplacingModuleFailsToLoad(t *testing.T) {
	m := module_tester.New(t)
	library := m.ReplacedLibrary()
	user := m.ReplacingUser(
		library,
		"package user\n\nimport \"other.test/lib/alfa\"\n\nfunc Use() string {\n\treturn alfa.SendMsg() + missing\n}\n",
	)
	o := m.Options(library, false, false, "./alfa/...")
	o.Replacing = []string{user}
	r := output.NewResultsWithDirectory(library)
	m.RunDefault(o, r)
	assert.String(
		t,
		"package alfa\n\nfunc SendMsg() string {\n\treturn \"x\"\n}\n",
		testutil.ReadFile(t, filepath.Join(library, "alfa", "send.go")),
	)
	assert.Integer(t, 1, len(fix_tester.Blocked(r.Entries)))
}
