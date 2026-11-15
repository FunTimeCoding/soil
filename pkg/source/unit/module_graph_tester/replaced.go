package module_graph_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func Replaced(t *testing.T) (string, string) {
	t.Helper()
	library := t.TempDir()
	system_tester.Write(
		t,
		library,
		"go.mod",
		"module other.test/lib\n\ngo 1.22\n",
	)
	system_tester.Write(
		t,
		library,
		"lib.go",
		"package lib\n\nfunc Name() string {\n\treturn \"lib\"\n}\n",
	)
	user := t.TempDir()
	system_tester.Write(
		t,
		user,
		"go.mod",
		fmt.Sprintf(
			"module example\n\ngo 1.22\n\nrequire other.test/lib v0.0.0\n\nreplace other.test/lib => %s\n",
			library,
		),
	)
	system_tester.Write(
		t,
		user,
		"user/user.go",
		"package user\n\nimport \"other.test/lib\"\n\nfunc Use() string {\n\treturn lib.Name()\n}\n",
	)

	return library, user
}
