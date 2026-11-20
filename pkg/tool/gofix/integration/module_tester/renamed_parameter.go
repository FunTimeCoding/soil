package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) RenamedParameter() string {
	o.t.Helper()

	return module.New(o.t,"example").
		File(
			"join.go",
			"package example\n\nfunc Join(\n\tmsg string,\n) string {\n\treturn msg\n}\n",
		).
		Directory()
}
