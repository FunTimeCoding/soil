package toolchain

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/fix_tester"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"testing"
)

func TestDefaultDiffShowsCombinedRun(t *testing.T) {
	m := module_tester.New(t)
	directory := m.Interacting()
	path := filepath.Join(directory, "interact.go")
	before := fix_tester.New(t).GoFiles(directory)
	printed := assert.Capture(
		t,
		func() {
			m.RunDefault(
				m.Options(directory, true, false, "./..."),
				output.NewResultsWithDirectory(directory),
			)
		},
	)
	assert.Any(t, before, fix_tester.New(t).GoFiles(directory))
	assert.String(
		t,
		fmt.Sprintf(
			"--- %[1]s\n+++ %[1]s\n package example\n \n-func Interact(dirName string, dirPath string) string {\n-\treturn join(dirName, dirPath, \"some-moderately-long-literal-value-here\")\n+func Interact(directoryName string, directoryPath string) string {\n+\treturn join(\n+\t\tdirectoryName,\n+\t\tdirectoryPath,\n+\t\t\"some-moderately-long-literal-value-here\",\n+\t)\n }\n \n",
			path,
		),
		printed,
	)
	m.RunDefault(
		m.Options(directory, false, false, "./..."),
		output.NewResultsWithDirectory(directory),
	)
	assert.String(
		t,
		"package example\n\nfunc Interact(directoryName string, directoryPath string) string {\n\treturn join(\n\t\tdirectoryName,\n\t\tdirectoryPath,\n\t\t\"some-moderately-long-literal-value-here\",\n\t)\n}\n\nfunc join(a, b, c string) string { return a + b + c }\n",
		fix_tester.New(t).GoFiles(directory)["interact.go"],
	)
}
