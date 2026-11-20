package module_graph_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
)

func (o *Tester) Replaced() (string, string) {
	o.t.Helper()
	library := o.t.TempDir()
	system_tester.New(o.t).Write(
		library,
		"go.mod",
		"module other.test/lib\n\ngo 1.22\n",
	)
	system_tester.New(o.t).Write(
		library,
		"lib.go",
		"package lib\n\nfunc Name() string {\n\treturn \"lib\"\n}\n",
	)
	user := o.t.TempDir()
	system_tester.New(o.t).Write(
		user,
		"go.mod",
		fmt.Sprintf(
			"module example\n\ngo 1.22\n\nrequire other.test/lib v0.0.0\n\nreplace other.test/lib => %s\n",
			library,
		),
	)
	system_tester.New(o.t).Write(
		user,
		"user/user.go",
		"package user\n\nimport \"other.test/lib\"\n\nfunc Use() string {\n\treturn lib.Name()\n}\n",
	)

	return library, user
}
