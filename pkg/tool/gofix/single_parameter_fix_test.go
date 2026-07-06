package gofix

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"path/filepath"
	"testing"
)

func TestSingleParameterFix(t *testing.T) {
	directory := writeSingleParameterTestModule(t)
	r := output.NewResultsWithDirectory(directory)
	runSingleParameterFixWithDirectory([]string{"./..."}, directory, r)
	t.Run(
		"MethodCollapsed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nimport \"context\"\n\ntype Client struct{}\n\nfunc (c *Client) Snapshot(x context.Context) (string, error) {\n\treturn \"\", nil\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "method.go"),
				),
			)
		},
	)
	t.Run(
		"FunctionCollapsed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc Process(name string) error {\n\treturn nil\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "function.go"),
				),
			)
		},
	)
	t.Run(
		"TwoParamsUntouched",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc TwoParams(\n\ta string,\n\tb string,\n) error {\n\treturn nil\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "two_params.go"),
				),
			)
		},
	)
	t.Run(
		"AlreadySingleLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc Short(x int) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "already_single.go"),
				),
			)
		},
	)
	t.Run(
		"TooLongUntouched",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc VeryLongFunctionNameThatWouldExceedTheLimit(\n\tparameterWithAVeryLongName string,\n) error {\n\treturn nil\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "too_long.go"),
				),
			)
		},
	)
	t.Run(
		"ResultEntries",
		func(t *testing.T) {
			applied := filterApplied(r.Entries)
			assert.Integer(t, 2, len(applied))
			assertResult(
				t,
				applied,
				"method.go",
				"collapsed single parameter (line 7)",
			)
			assertResult(
				t,
				applied,
				"function.go",
				"collapsed single parameter (line 3)",
			)
		},
	)
}

func TestSingleParameterFixWithTestFiles(t *testing.T) {
	directory := writeSingleParameterTestModuleWithTests(t)
	r := output.NewResultsWithDirectory(directory)
	runSingleParameterFixWithDirectory([]string{"./..."}, directory, r)
	t.Run(
		"NotCorrupted",
		func(t *testing.T) {
			assert.String(
				t,
				"package tested\n\nfunc FindLatest(v []string) *string {\n\treturn nil\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "find.go"),
				),
			)
		},
	)
}

func writeSingleParameterTestModuleWithTests(t *testing.T) string {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteFile(
		t,
		directory,
		"go.mod",
		"module example\n\ngo 1.22\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"find.go",
		"package tested\n\nfunc FindLatest(\n\tv []string,\n) *string {\n\treturn nil\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"find_test.go",
		"package tested\n\nimport \"testing\"\n\nfunc TestFindLatest(t *testing.T) {}\n",
	)

	return directory
}

func writeSingleParameterTestModule(t *testing.T) string {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteFile(
		t,
		directory,
		"go.mod",
		"module example\n\ngo 1.22\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"method.go",
		"package example\n\nimport \"context\"\n\ntype Client struct{}\n\nfunc (c *Client) Snapshot(\n\tx context.Context,\n) (string, error) {\n\treturn \"\", nil\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"function.go",
		"package example\n\nfunc Process(\n\tname string,\n) error {\n\treturn nil\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"two_params.go",
		"package example\n\nfunc TwoParams(\n\ta string,\n\tb string,\n) error {\n\treturn nil\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"already_single.go",
		"package example\n\nfunc Short(x int) {}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"too_long.go",
		"package example\n\nfunc VeryLongFunctionNameThatWouldExceedTheLimit(\n\tparameterWithAVeryLongName string,\n) error {\n\treturn nil\n}\n",
	)

	return directory
}
