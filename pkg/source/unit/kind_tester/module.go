package kind_tester

import "github.com/funtimecoding/soil/pkg/system/system_tester"

func (o *Tester) Module() string {
	o.t.Helper()
	directory := o.t.TempDir()
	system_tester.New(o.t).Write(
		directory,
		"go.mod",
		"module example\n\ngo 1.22\n",
	)
	system_tester.New(o.t).Write(
		directory,
		"one/one.go",
		"package one\n\nimport \"strings\"\n\nfunc One() string {\n\treturn strings.ToUpper(\"a\")\n}\n",
	)
	system_tester.New(o.t).Write(
		directory,
		"two/two.go",
		"package two\n\nimport \"example/one\"\n\nfunc Two() string {\n\treturn one.One()\n}\n",
	)

	return directory
}
