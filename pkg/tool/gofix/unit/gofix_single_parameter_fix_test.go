package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/unit/module_tester"
	"path/filepath"
	"testing"
)

func TestSingleParameterFix(t *testing.T) {
	directory := module_tester.SingleParameter(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunSingleParameterFixWithDirectory(
		[]string{"./..."},
		directory,
		false,
		r,
	)
	t.Run(
		"MethodCollapsed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nimport \"context\"\n\ntype Client struct{}\n\nfunc (c *Client) Snapshot(x context.Context) (string, error) {\n\treturn \"\", nil\n}\n",
				testutil.ReadFile(t, filepath.Join(directory, "method.go")),
			)
		},
	)
	t.Run(
		"FunctionCollapsed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc Process(name string) error {\n\treturn nil\n}\n",
				testutil.ReadFile(t, filepath.Join(directory, "function.go")),
			)
		},
	)
	t.Run(
		"TwoParamsUntouched",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc TwoParams(\n\ta string,\n\tb string,\n) error {\n\treturn nil\n}\n",
				testutil.ReadFile(t, filepath.Join(directory, "two_params.go")),
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
				testutil.ReadFile(t, filepath.Join(directory, "too_long.go")),
			)
		},
	)
	t.Run(
		"ResultEntries",
		func(t *testing.T) {
			applied := filterApplied(r.Entries)
			assert.Integer(t, 2, len(applied))
			assertResultAt(
				t,
				applied,
				"method.go",
				7,
				"collapsed single parameter",
			)
			assertResultAt(
				t,
				applied,
				"function.go",
				3,
				"collapsed single parameter",
			)
		},
	)
}

func TestSingleParameterFixWithTestFiles(t *testing.T) {
	directory := module_tester.SingleParameterWithTests(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunSingleParameterFixWithDirectory(
		[]string{"./..."},
		directory,
		false,
		r,
	)
	t.Run(
		"NotCorrupted",
		func(t *testing.T) {
			assert.String(
				t,
				"package tested\n\nfunc FindLatest(v []string) *string {\n\treturn nil\n}\n",
				testutil.ReadFile(t, filepath.Join(directory, "find.go")),
			)
		},
	)
}
