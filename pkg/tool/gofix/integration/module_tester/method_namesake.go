package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) MethodNamesake() string {
	o.t.Helper()

	return module.New(o.t,"testmodule").
		File(
			"pkg/helper/helper.go",
			"package helper\n\nfunc Value() string { return \"hello\" }\n",
		).
		File(
			"holder.go",
			"package example\n\nimport h \"testmodule/pkg/helper\"\n\ntype Holder struct{}\n\nfunc (o *Holder) helper() string {\n\treturn h.Value()\n}\n",
		).
		Directory()
}
