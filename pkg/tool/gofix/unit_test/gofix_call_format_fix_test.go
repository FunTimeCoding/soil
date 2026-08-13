package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"path/filepath"
	"testing"
)

func TestCallFormatFix(t *testing.T) {
	directory := writeCallFormatTestModule(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunCallFormatFixWithDirectory([]string{"./..."}, directory, r)
	t.Run(
		"LongSingleLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc LongSingleLine() {\n\ttwoArgs(\n\t\t\"something-long-enough\",\n\t\t\"to-push-this-well-past-the-eighty-character-column-limit\",\n\t)\n}\n\nfunc twoArgs(a, b string) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "long_single_line.go"),
				),
			)
		},
	)
	t.Run(
		"SharedLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\ntype Options struct {\n\tValue int\n}\n\nfunc SharedLine() {\n\twithStruct(\n\t\t\"name\",\n\t\tOptions{Value: 1},\n\t)\n}\n\nfunc withStruct(a string, b Options) {}\n",
				testutil.ReadFile(t, filepath.Join(directory, "shared_line.go")),
			)
		},
	)
	t.Run(
		"FirstArgOnParenLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc FirstArgOnParenLine() {\n\twithMap(\n\t\t\"name\",\n\t\tmap[string]any{\"key\": \"value\"},\n\t)\n}\n\nfunc withMap(a string, b map[string]any) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "first_arg_on_paren_line.go"),
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
		"CompliantCollapsed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc Compliant() {\n\ttwoArgs(\"alpha\", \"bravo\")\n\ttwoArgs(\"alpha\", \"bravo\")\n}\n",
				testutil.ReadFile(t, filepath.Join(directory, "compliant.go")),
			)
		},
	)
	t.Run(
		"DeepMethodCall",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\ntype Logger struct{}\n\nfunc (l *Logger) Structured(args ...string) {}\n\ntype Poller struct {\n\tlogger *Logger\n}\n\nfunc (p *Poller) Run() {\n\tdefer func() {\n\t\tif v := recover(); v != nil {\n\t\t\tp.logger.Structured(\"recover failed\", \"error\", \"value\")\n\t\t}\n\t}()\n}\n",
				testutil.ReadFile(t, filepath.Join(directory, "deep_method.go")),
			)
		},
	)
	t.Run(
		"MultipleViolationsInOneCall",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc MultipleViolations() {\n\tfourArgs(\"adopted\", \"bravo\", \"charlie\", \"delta\")\n}\n\nfunc fourArgs(a, b, c, d string) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "multiple_violations.go"),
				),
			)
		},
	)
	t.Run(
		"CollapseMultiLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc CollapseMultiLine() {\n\ttwoArgs(\"alpha\", \"bravo\")\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "collapse_multi_line.go"),
				),
			)
		},
	)
	t.Run(
		"CollapseSingleArg",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc CollapseSingleArg() {\n\toneArg(\"short\")\n}\n\nfunc oneArg(a string) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "collapse_single_arg.go"),
				),
			)
		},
	)
	t.Run(
		"BoundaryAt80Collapses",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc BoundaryAt80() {\n\ttwoArgs(\"aaaaaaaaaaaaaaaaaaaaaa\", \"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb\")\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "boundary_at_80.go"),
				),
			)
		},
	)
	t.Run(
		"BoundaryAt81Stays",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc BoundaryAt81() {\n\ttwoArgs(\n\t\t\"aaaaaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t)\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "boundary_at_81.go"),
				),
			)
		},
	)
	t.Run(
		"StructFieldPaddingPreventsCollapse",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\ntype Thing struct {\n\tLongFieldName string\n\tShort         string\n}\n\nfunc StructFieldPadding() Thing {\n\treturn Thing{\n\t\tLongFieldName: \"value\",\n\t\tShort: someFunc(\n\t\t\t\"aaaaaaaaaaaaaaaaa\",\n\t\t\t\"bbbbbbbbbbbbbbbbbbbbbbbb\",\n\t\t),\n\t}\n}\n\nfunc someFunc(a, b string) string { return a }\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "struct_field_padding.go"),
				),
			)
		},
	)
	t.Run(
		"TrailingContentPreventsCollapse",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\ntype Chain struct{}\n\nfunc (c Chain) Method(a, b string) Chain { return c }\nfunc (c Chain) Suffix() {}\n\nfunc TrailingContent() {\n\tChain{}.Method(\n\t\t\"aaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t).Suffix()\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "trailing_content.go"),
				),
			)
		},
	)
	t.Run(
		"EllipsisStaysMultiLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc EllipsisStays(items []string) {\n\tvariadicFunc(\n\t\titems...,\n\t)\n}\n\nfunc variadicFunc(args ...string) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "ellipsis_stays.go"),
				),
			)
		},
	)
	t.Run(
		"VarBlockPaddingPreventsCollapse",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nvar (\n\tLongName = someFunc(\"short\", \"args\")\n\tX        = someFunc(\n\t\t\"aaaaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t)\n)\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "var_block_padding.go"),
				),
			)
		},
	)
	t.Run(
		"CollapsePreservesLong",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc CollapsePreservesLong() {\n\ttwoArgs(\n\t\t\"something-long-enough\",\n\t\t\"to-push-this-well-past-the-eighty-character-column-limit\",\n\t)\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "collapse_preserves_long.go"),
				),
			)
		},
	)
	t.Run(
		"ResultEntries",
		func(t *testing.T) {
			applied := filterApplied(r.Entries)
			assert.Integer(t, 12, len(applied))
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
			assertResult(
				t,
				applied,
				"collapse_multi_line.go",
				"formatted call (line 4)",
			)
			assertResult(t, applied, "compliant.go", "formatted call (line 5)")
			assertResult(
				t,
				applied,
				"collapse_single_arg.go",
				"formatted call (line 4)",
			)
			assertResult(
				t,
				applied,
				"boundary_at_80.go",
				"formatted call (line 4)",
			)
		},
	)
}

func writeCallFormatTestModule(t *testing.T) string {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteFile(t, directory, "go.mod", "module example\n\ngo 1.22\n")
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
	testutil.WriteFile(
		t,
		directory,
		"collapse_single_arg.go",
		"package example\n\nfunc CollapseSingleArg() {\n\toneArg(\n\t\t\"short\",\n\t)\n}\n\nfunc oneArg(a string) {}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"boundary_at_80.go",
		"package example\n\nfunc BoundaryAt80() {\n\ttwoArgs(\n\t\t\"aaaaaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t)\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"boundary_at_81.go",
		"package example\n\nfunc BoundaryAt81() {\n\ttwoArgs(\n\t\t\"aaaaaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t)\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"collapse_multi_line.go",
		"package example\n\nfunc CollapseMultiLine() {\n\ttwoArgs(\n\t\t\"alpha\",\n\t\t\"bravo\",\n\t)\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"collapse_preserves_long.go",
		"package example\n\nfunc CollapsePreservesLong() {\n\ttwoArgs(\n\t\t\"something-long-enough\",\n\t\t\"to-push-this-well-past-the-eighty-character-column-limit\",\n\t)\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"struct_field_padding.go",
		"package example\n\ntype Thing struct {\n\tLongFieldName string\n\tShort         string\n}\n\nfunc StructFieldPadding() Thing {\n\treturn Thing{\n\t\tLongFieldName: \"value\",\n\t\tShort: someFunc(\n\t\t\t\"aaaaaaaaaaaaaaaaa\",\n\t\t\t\"bbbbbbbbbbbbbbbbbbbbbbbb\",\n\t\t),\n\t}\n}\n\nfunc someFunc(a, b string) string { return a }\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"trailing_content.go",
		"package example\n\ntype Chain struct{}\n\nfunc (c Chain) Method(a, b string) Chain { return c }\nfunc (c Chain) Suffix() {}\n\nfunc TrailingContent() {\n\tChain{}.Method(\n\t\t\"aaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t).Suffix()\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"ellipsis_stays.go",
		"package example\n\nfunc EllipsisStays(items []string) {\n\tvariadicFunc(\n\t\titems...,\n\t)\n}\n\nfunc variadicFunc(args ...string) {}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"var_block_padding.go",
		"package example\n\nvar (\n\tLongName = someFunc(\"short\", \"args\")\n\tX        = someFunc(\n\t\t\"aaaaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t)\n)\n",
	)

	return directory
}
