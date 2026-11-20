package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) Interacting() string {
	o.t.Helper()

	return module.New(o.t,"example").
		File(
			"interact.go",
			"package example\n\nfunc Interact(dirName string, dirPath string) string {\n\treturn join(dirName, dirPath, \"some-moderately-long-literal-value-here\")\n}\n\nfunc join(a, b, c string) string { return a + b + c }\n",
		).
		Directory()
}
