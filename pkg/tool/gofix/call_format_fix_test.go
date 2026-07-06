package gofix

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"path/filepath"
	"testing"
)

func TestCallFormatFix(t *testing.T) {
	directory := writeCallFormatTestModule(t)
	r := output.NewResultsWithDirectory(directory)
	runCallFormatFixWithDirectory([]string{"./..."}, directory, r)
	t.Run(
		"LongSingleLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc LongSingleLine() {\n\ttwoArgs(\n\t\t\"something-long-enough\",\n\t\t\"to-push-this-well-past-the-eighty-character-column-limit\",\n\t)\n}\n\nfunc twoArgs(a, b string) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(
						directory,
						"long_single_line.go",
					),
				),
			)
		},
	)
	t.Run(
		"SharedLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\ntype Options struct {\n\tValue int\n}\n\nfunc SharedLine() {\n\twithStruct(\n\t\t\"name\",\n\t\tOptions{\n\t\t\tValue: 1,\n\t\t},\n\t)\n}\n\nfunc withStruct(a string, b Options) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "shared_line.go"),
				),
			)
		},
	)
	t.Run(
		"FirstArgOnParenLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc FirstArgOnParenLine() {\n\twithMap(\n\t\t\"name\",\n\t\tmap[string]any{\n\t\t\t\"key\": \"value\",\n\t\t},\n\t)\n}\n\nfunc withMap(a string, b map[string]any) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(
						directory,
						"first_arg_on_paren_line.go",
					),
				),
			)
		},
	)
	t.Run(
		"NestedIndent",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc NestedIndent() {\n\tif true {\n\t\tif true {\n\t\t\ttwoArgs(\n\t\t\t\t\"something-long-enough-to-exceed\",\n\t\t\t\t\"the-eighty-character-limit-at-this-indent-level\",\n\t\t\t)\n\t\t}\n\t}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "nested_indent.go"),
				),
			)
		},
	)
	t.Run(
		"CompliantUntouched",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc Compliant() {\n\ttwoArgs(\"alpha\", \"bravo\")\n\ttwoArgs(\n\t\t\"alpha\",\n\t\t\"bravo\",\n\t)\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "compliant.go"),
				),
			)
		},
	)
	t.Run(
		"DeepMethodCall",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\ntype Logger struct{}\n\nfunc (l *Logger) Structured(args ...string) {}\n\ntype Poller struct {\n\tlogger *Logger\n}\n\nfunc (p *Poller) Run() {\n\tdefer func() {\n\t\tif v := recover(); v != nil {\n\t\t\tp.logger.Structured(\n\t\t\t\t\"recover failed\",\n\t\t\t\t\"error\",\n\t\t\t\t\"value\",\n\t\t\t)\n\t\t}\n\t}()\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "deep_method.go"),
				),
			)
		},
	)
	t.Run(
		"MultipleViolationsInOneCall",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc MultipleViolations() {\n\tfourArgs(\n\t\t\"adopted\",\n\t\t\"bravo\",\n\t\t\"charlie\",\n\t\t\"delta\",\n\t)\n}\n\nfunc fourArgs(a, b, c, d string) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(
						directory,
						"multiple_violations.go",
					),
				),
			)
		},
	)
	t.Run(
		"ResultEntries",
		func(t *testing.T) {
			applied := filterApplied(r.Entries)
			assert.Integer(t, 6, len(applied))
			assertResult(
				t,
				applied,
				"long_single_line.go",
				"formatted call (line 4)",
			)
			assertResult(
				t,
				applied,
				"shared_line.go",
				"formatted call (line 8)",
			)
			assertResult(
				t,
				applied,
				"first_arg_on_paren_line.go",
				"formatted call (line 4)",
			)
			assertResult(
				t,
				applied,
				"nested_indent.go",
				"formatted call (line 6)",
			)
			assertResult(
				t,
				applied,
				"deep_method.go",
				"formatted call (line 14)",
			)
			assertResult(
				t,
				applied,
				"multiple_violations.go",
				"formatted call (line 4)",
			)
		},
	)
}

func writeCallFormatTestModule(t *testing.T) string {
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
		"long_single_line.go",
		"package example\n\nfunc LongSingleLine() {\n\ttwoArgs(\"something-long-enough\", \"to-push-this-well-past-the-eighty-character-column-limit\")\n}\n\nfunc twoArgs(a, b string) {}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"shared_line.go",
		"package example\n\ntype Options struct {\n\tValue int\n}\n\nfunc SharedLine() {\n\twithStruct(\n\t\t\"name\", Options{\n\t\t\tValue: 1,\n\t\t},\n\t)\n}\n\nfunc withStruct(a string, b Options) {}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"first_arg_on_paren_line.go",
		"package example\n\nfunc FirstArgOnParenLine() {\n\twithMap(\"name\",\n\t\tmap[string]any{\n\t\t\t\"key\": \"value\",\n\t\t},\n\t)\n}\n\nfunc withMap(a string, b map[string]any) {}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"nested_indent.go",
		"package example\n\nfunc NestedIndent() {\n\tif true {\n\t\tif true {\n\t\t\ttwoArgs(\"something-long-enough-to-exceed\", \"the-eighty-character-limit-at-this-indent-level\")\n\t\t}\n\t}\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"compliant.go",
		"package example\n\nfunc Compliant() {\n\ttwoArgs(\"alpha\", \"bravo\")\n\ttwoArgs(\n\t\t\"alpha\",\n\t\t\"bravo\",\n\t)\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"multiple_violations.go",
		"package example\n\nfunc MultipleViolations() {\n\tfourArgs(\n\t\t\"adopted\", \"bravo\",\n\t\t\"charlie\", \"delta\",\n\t)\n}\n\nfunc fourArgs(a, b, c, d string) {}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"deep_method.go",
		"package example\n\ntype Logger struct{}\n\nfunc (l *Logger) Structured(args ...string) {}\n\ntype Poller struct {\n\tlogger *Logger\n}\n\nfunc (p *Poller) Run() {\n\tdefer func() {\n\t\tif v := recover(); v != nil {\n\t\t\tp.logger.Structured(\n\t\t\t\t\"recover failed\",\n\t\t\t\t\"error\", \"value\",\n\t\t\t)\n\t\t}\n\t}()\n}\n",
	)

	return directory
}
