package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/fix_tester"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"testing"
)

func TestSingleParameterFix(t *testing.T) {
	directory := module_tester.New(t).SingleParameter()
	r := output.NewResultsWithDirectory(directory)
	gofix.RunSingleParameterFix([]string{"./..."}, directory, false, r)
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
		"ExpandedFunctionTypeCollapsed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc Expanded(f func(int, []byte) error) error {\n\treturn nil\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "expanded_function_type.go"),
				),
			)
		},
	)
	t.Run(
		"HalfExpandedFunctionTypeCollapsed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc HalfExpanded(f func(int, []byte) error) error {\n\treturn nil\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "half_expanded_function_type.go"),
				),
			)
		},
	)
	t.Run(
		"CommentedFunctionTypeUntouched",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc Commented(\n\tf func(\n\t\tint, // status\n\t\t[]byte,\n\t) error,\n) error {\n\treturn nil\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "commented_function_type.go"),
				),
			)
		},
	)
	t.Run(
		"StructTypeUntouched",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc Structured(\n\tv struct {\n\t\tA int\n\t\tB int\n\t},\n) error {\n\treturn nil\n}\n",
				testutil.ReadFile(t, filepath.Join(directory, "struct_type.go")),
			)
		},
	)
	t.Run(
		"ResultEntries",
		func(t *testing.T) {
			applied := fix_tester.Applied(r.Entries)
			assert.Integer(t, 4, len(applied))
			fix_tester.New(t).LineResult(
				applied,
				"method.go",
				7,
				"collapsed single parameter",
			)
			fix_tester.New(t).LineResult(
				applied,
				"function.go",
				3,
				"collapsed single parameter",
			)
			fix_tester.New(t).LineResult(
				applied,
				"expanded_function_type.go",
				3,
				"collapsed single parameter",
			)
			fix_tester.New(t).LineResult(
				applied,
				"half_expanded_function_type.go",
				3,
				"collapsed single parameter",
			)
		},
	)
}

func TestSingleParameterFixSettles(t *testing.T) {
	directory := module_tester.New(t).SingleParameter()
	gofix.RunSingleParameterFix(
		[]string{"./..."},
		directory,
		false,
		output.NewResultsWithDirectory(directory),
	)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunSingleParameterFix([]string{"./..."}, directory, false, r)
	assert.Integer(t, 0, len(fix_tester.Applied(r.Entries)))
}

func TestSingleParameterFixWithTestFiles(t *testing.T) {
	directory := module_tester.New(t).SingleParameterWithTests()
	r := output.NewResultsWithDirectory(directory)
	gofix.RunSingleParameterFix([]string{"./..."}, directory, false, r)
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
