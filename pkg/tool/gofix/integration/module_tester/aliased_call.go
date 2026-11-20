package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) AliasedCall() string {
	o.t.Helper()

	return module.New(o.t,"testmodule").
		File(
			"pkg/helper/helper.go",
			"package helper\n\nfunc Join(a, b, c string) string { return a + b + c }\n",
		).
		File(
			"use.go",
			"package example\n\nimport longerHelperAlias \"testmodule/pkg/helper\"\n\nfunc Use() string {\n\treturn longerHelperAlias.Join(\n\t\t\"first-literal-value\",\n\t\t\"second-literal-value\",\n\t\t\"third\",\n\t)\n}\n",
		).
		Directory()
}
