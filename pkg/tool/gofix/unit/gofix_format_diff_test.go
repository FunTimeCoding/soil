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

func TestFormatDiffShowsTheFullRun(t *testing.T) {
	directory := module_tester.FormatDiff(t)
	path := filepath.Join(directory, "exploded_nested.go")
	before := readGoFiles(t, directory)
	printed := assert.Capture(
		t,
		func() {
			gofix.RunFormatFixWithDirectory(
				[]string{"./..."},
				directory,
				true,
				output.NewResultsWithDirectory(directory),
			)
		},
	)
	assert.Any(t, before, readGoFiles(t, directory))
	assert.String(
		t,
		fmt.Sprintf(
			"--- %[1]s\n+++ %[1]s\n@@\n \n func ExplodedNested(value string) string {\n-\treturn outerOne(\n-\t\tinnerOne(\n-\t\t\tvalue,\n-\t\t),\n-\t)\n+\treturn outerOne(innerOne(value))\n }\n \n",
			path,
		),
		printed,
	)
	gofix.RunFormatFixWithDirectory(
		[]string{"./..."},
		directory,
		false,
		output.NewResultsWithDirectory(directory),
	)
	assert.String(
		t,
		"package example\n\nfunc ExplodedNested(value string) string {\n\treturn outerOne(innerOne(value))\n}\n\nfunc outerOne(a string) string { return a }\n\nfunc innerOne(a string) string { return a }\n",
		testutil.ReadFile(t, path),
	)
}
