package module_tester

import "testing"

func CallFormat(t *testing.T) string {
	t.Helper()

	return New(t, "example").
		File(
			"long_single_line.go",
			"package example\n\nfunc LongSingleLine() {\n\ttwoArgs(\"something-long-enough\", \"to-push-this-well-past-the-eighty-character-column-limit\")\n}\n\nfunc twoArgs(a, b string) {}\n",
		).
		File(
			"shared_line.go",
			"package example\n\ntype Options struct {\n\tValue int\n}\n\nfunc SharedLine() {\n\twithStruct(\n\t\t\"name\", Options{\n\t\t\tValue: 1,\n\t\t},\n\t)\n}\n\nfunc withStruct(a string, b Options) {}\n",
		).
		File(
			"first_arg_on_paren_line.go",
			"package example\n\nfunc FirstArgOnParenLine() {\n\twithMap(\"name\",\n\t\tmap[string]any{\n\t\t\t\"key\": \"value\",\n\t\t},\n\t)\n}\n\nfunc withMap(a string, b map[string]any) {}\n",
		).
		File(
			"nested_indent.go",
			"package example\n\nfunc NestedIndent() {\n\tif true {\n\t\tif true {\n\t\t\ttwoArgs(\"something-long-enough-to-exceed\", \"the-eighty-character-limit-at-this-indent-level\")\n\t\t}\n\t}\n}\n",
		).
		File(
			"compliant.go",
			"package example\n\nfunc Compliant() {\n\ttwoArgs(\"alfa\", \"bravo\")\n\ttwoArgs(\n\t\t\"alfa\",\n\t\t\"bravo\",\n\t)\n}\n",
		).
		File(
			"multiple_violations.go",
			"package example\n\nfunc MultipleViolations() {\n\tfourArgs(\n\t\t\"adopted\", \"bravo\",\n\t\t\"charlie\", \"delta\",\n\t)\n}\n\nfunc fourArgs(a, b, c, d string) {}\n",
		).
		File(
			"deep_method.go",
			"package example\n\ntype Logger struct{}\n\nfunc (l *Logger) Structured(args ...string) {}\n\ntype Poller struct {\n\tlogger *Logger\n}\n\nfunc (p *Poller) Run() {\n\tdefer func() {\n\t\tif v := recover(); v != nil {\n\t\t\tp.logger.Structured(\n\t\t\t\t\"recover failed\",\n\t\t\t\t\"error\", \"value\",\n\t\t\t)\n\t\t}\n\t}()\n}\n",
		).
		File(
			"collapse_single_arg.go",
			"package example\n\nfunc CollapseSingleArg() {\n\toneArg(\n\t\t\"short\",\n\t)\n}\n\nfunc oneArg(a string) {}\n",
		).
		File(
			"boundary_at_80.go",
			"package example\n\nfunc BoundaryAt80() {\n\ttwoArgs(\n\t\t\"aaaaaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t)\n}\n",
		).
		File(
			"boundary_at_81.go",
			"package example\n\nfunc BoundaryAt81() {\n\ttwoArgs(\n\t\t\"aaaaaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t)\n}\n",
		).
		File(
			"collapse_multi_line.go",
			"package example\n\nfunc CollapseMultiLine() {\n\ttwoArgs(\n\t\t\"alfa\",\n\t\t\"bravo\",\n\t)\n}\n",
		).
		File(
			"collapse_preserves_long.go",
			"package example\n\nfunc CollapsePreservesLong() {\n\ttwoArgs(\n\t\t\"something-long-enough\",\n\t\t\"to-push-this-well-past-the-eighty-character-column-limit\",\n\t)\n}\n",
		).
		File(
			"struct_field_padding.go",
			"package example\n\ntype Thing struct {\n\tLongFieldName string\n\tShort         string\n}\n\nfunc StructFieldPadding() Thing {\n\treturn Thing{\n\t\tLongFieldName: \"value\",\n\t\tShort: someFunc(\n\t\t\t\"aaaaaaaaaaaaaaaaa\",\n\t\t\t\"bbbbbbbbbbbbbbbbbbbbbbbb\",\n\t\t),\n\t}\n}\n\nfunc someFunc(a, b string) string { return a }\n",
		).
		File(
			"trailing_content.go",
			"package example\n\ntype Chain struct{}\n\nfunc (c Chain) Method(a, b string) Chain { return c }\nfunc (c Chain) Suffix() {}\n\nfunc TrailingContent() {\n\tChain{}.Method(\n\t\t\"aaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t).Suffix()\n}\n",
		).
		File(
			"ellipsis_stays.go",
			"package example\n\nfunc EllipsisStays(items []string) {\n\tvariadicFunc(\n\t\titems...,\n\t)\n}\n\nfunc variadicFunc(args ...string) {}\n",
		).
		File(
			"var_block_padding.go",
			"package example\n\nvar (\n\tLongName = someFunc(\"short\", \"args\")\n\tX        = someFunc(\n\t\t\"aaaaaaaaaaaaaaaaaaaaa\",\n\t\t\"bbbbbbbbbbbbbbbbbbbbbbbbbbbbb\",\n\t)\n)\n",
		).
		Directory()
}
