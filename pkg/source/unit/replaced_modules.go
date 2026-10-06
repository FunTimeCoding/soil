package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func replacedModules(t *testing.T) (string, string) {
	t.Helper()
	library := t.TempDir()
	testutil.WriteFile(
		t,
		library,
		"go.mod",
		"module other.test/lib\n\ngo 1.22\n",
	)
	testutil.WriteFile(
		t,
		library,
		"lib.go",
		"package lib\n\nfunc Name() string {\n\treturn \"lib\"\n}\n",
	)
	user := t.TempDir()
	testutil.WriteFile(
		t,
		user,
		"go.mod",
		fmt.Sprintf(
			"module example\n\ngo 1.22\n\nrequire other.test/lib v0.0.0\n\nreplace other.test/lib => %s\n",
			library,
		),
	)
	testutil.WriteFile(
		t,
		user,
		"user/user.go",
		"package user\n\nimport \"other.test/lib\"\n\nfunc Use() string {\n\treturn lib.Name()\n}\n",
	)

	return library, user
}
