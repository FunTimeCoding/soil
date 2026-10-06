package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/unit/module_tester"
	"path/filepath"
	"testing"
)

func TestDefaultDiffShowsTheCombinedRun(t *testing.T) {
	directory := module_tester.Interacting(t)
	path := filepath.Join(directory, "interact.go")
	before := readGoFiles(t, directory)
	printed := assert.Capture(
		t,
		func() {
			gofix.RunDefault(
				module_tester.Options(t, directory, true, false, "./..."),
				output.NewResultsWithDirectory(directory),
			)
		},
	)
	assert.Any(t, before, readGoFiles(t, directory))
	assert.String(
		t,
		fmt.Sprintf(
			"--- %[1]s\n+++ %[1]s\n package example\n \n-func Interact(dirName string, dirPath string) string {\n-\treturn join(dirName, dirPath, \"some-moderately-long-literal-value-here\")\n+func Interact(directoryName string, directoryPath string) string {\n+\treturn join(\n+\t\tdirectoryName,\n+\t\tdirectoryPath,\n+\t\t\"some-moderately-long-literal-value-here\",\n+\t)\n }\n \n",
			path,
		),
		printed,
	)
	gofix.RunDefault(
		module_tester.Options(t, directory, false, false, "./..."),
		output.NewResultsWithDirectory(directory),
	)
	assert.String(
		t,
		"package example\n\nfunc Interact(directoryName string, directoryPath string) string {\n\treturn join(\n\t\tdirectoryName,\n\t\tdirectoryPath,\n\t\t\"some-moderately-long-literal-value-here\",\n\t)\n}\n\nfunc join(a, b, c string) string { return a + b + c }\n",
		readGoFiles(t, directory)["interact.go"],
	)
}
