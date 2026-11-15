package index_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func KeyedModule(t *testing.T) string {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteModFile(t, directory, "example")
	system_tester.Write(
		t,
		directory,
		"charlie/charlie.go",
		"package charlie\n\nfunc Charlie() int {\n\treturn 1\n}\n",
	)
	system_tester.Write(
		t,
		directory,
		"bravo/bravo.go",
		"package bravo\n\nimport \"example/charlie\"\n\nfunc Bravo() int {\n\treturn charlie.Charlie()\n}\n",
	)
	system_tester.Write(
		t,
		directory,
		"alfa/alfa.go",
		"package alfa\n\nimport \"example/bravo\"\n\nfunc Alfa() int {\n\treturn bravo.Bravo()\n}\n",
	)
	system_tester.Write(
		t,
		directory,
		"alfa/alfa_test.go",
		"package alfa\n\nimport \"testing\"\n\nfunc TestAlfa(t *testing.T) {}\n",
	)
	system_tester.Write(
		t,
		directory,
		"delta/delta.go",
		"package delta\n\nfunc Delta() int {\n\treturn 4\n}\n",
	)

	return directory
}
