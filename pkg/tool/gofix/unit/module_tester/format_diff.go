package module_tester

import "testing"

func FormatDiff(t *testing.T) string {
	t.Helper()

	return New(t, "example").
		File(
			"exploded_nested.go",
			"package example\n\nfunc ExplodedNested(value string) string {\n\treturn outerOne(\n\t\tinnerOne(\n\t\t\tvalue,\n\t\t),\n\t)\n}\n\nfunc outerOne(a string) string { return a }\n\nfunc innerOne(a string) string { return a }\n",
		).
		Directory()
}
