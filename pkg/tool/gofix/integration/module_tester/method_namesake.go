package module_tester

import "testing"

func MethodNamesake(t *testing.T) string {
	t.Helper()

	return New(t, "testmodule").
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
