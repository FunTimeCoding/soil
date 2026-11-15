package kind_tester

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func Module(t *testing.T) string {
	t.Helper()
	directory := t.TempDir()
	system_tester.Write(t, directory, "go.mod", "module example\n\ngo 1.22\n")
	system_tester.Write(
		t,
		directory,
		"one/one.go",
		"package one\n\nimport \"strings\"\n\nfunc One() string {\n\treturn strings.ToUpper(\"a\")\n}\n",
	)
	system_tester.Write(
		t,
		directory,
		"two/two.go",
		"package two\n\nimport \"example/one\"\n\nfunc Two() string {\n\treturn one.One()\n}\n",
	)

	return directory
}
