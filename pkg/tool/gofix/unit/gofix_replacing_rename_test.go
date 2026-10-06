package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/unit/module_tester"
	"path/filepath"
	"testing"
)

func TestExportedRenameReachesAReplacingModule(t *testing.T) {
	for _, full := range []bool{false, true} {
		library := module_tester.ReplacedLibrary(t)
		user := module_tester.ReplacingUser(
			t,
			library,
			"package user\n\nimport \"other.test/lib/alfa\"\n\nfunc Use() string {\n\treturn alfa.SendMsg()\n}\n",
		)
		o := module_tester.Options(t, library, false, full, "./alfa/...")
		o.Replacing = []string{user}
		gofix.RunDefault(o, output.NewResultsWithDirectory(library))
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

func TestExportedRenameDiffLeavesAReplacingModule(t *testing.T) {
	library := module_tester.ReplacedLibrary(t)
	source := "package user\n\nimport \"other.test/lib/alfa\"\n\nfunc Use() string {\n\treturn alfa.SendMsg()\n}\n"
	user := module_tester.ReplacingUser(t, library, source)
	o := module_tester.Options(t, library, true, false, "./alfa/...")
	o.Replacing = []string{user}
	printed := assert.Capture(
		t,
		func() {
			gofix.RunDefault(o, output.NewResultsWithDirectory(library))
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

func TestExportedRenameRefusedWhenAReplacingModuleFailsToLoad(t *testing.T) {
	library := module_tester.ReplacedLibrary(t)
	user := module_tester.ReplacingUser(
		t,
		library,
		"package user\n\nimport \"other.test/lib/alfa\"\n\nfunc Use() string {\n\treturn alfa.SendMsg() + missing\n}\n",
	)
	o := module_tester.Options(t, library, false, false, "./alfa/...")
	o.Replacing = []string{user}
	r := output.NewResultsWithDirectory(library)
	gofix.RunDefault(o, r)
	assert.String(
		t,
		"package alfa\n\nfunc SendMsg() string {\n\treturn \"x\"\n}\n",
		testutil.ReadFile(t, filepath.Join(library, "alfa", "send.go")),
	)
	assert.Integer(t, 1, len(filterBlocked(r.Entries)))
}
