package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func kindModule(t *testing.T) string {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteFile(t, directory, "go.mod", "module example\n\ngo 1.22\n")
	testutil.WriteFile(
		t,
		directory,
		"one/one.go",
		"package one\n\nimport \"strings\"\n\nfunc One() string {\n\treturn strings.ToUpper(\"a\")\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"two/two.go",
		"package two\n\nimport \"example/one\"\n\nfunc Two() string {\n\treturn one.One()\n}\n",
	)

	return directory
}
