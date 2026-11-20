package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) FormatDiff() string {
	o.t.Helper()

	return module.New(o.t,"example").
		File(
			"exploded_nested.go",
			"package example\n\nfunc ExplodedNested(value string) string {\n\treturn outerOne(\n\t\tinnerOne(\n\t\t\tvalue,\n\t\t),\n\t)\n}\n\nfunc outerOne(a string) string { return a }\n\nfunc innerOne(a string) string { return a }\n",
		).
		Directory()
}
