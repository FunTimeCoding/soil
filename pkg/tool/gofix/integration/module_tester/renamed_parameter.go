package module_tester

import "testing"

func RenamedParameter(t *testing.T) string {
	t.Helper()

	return New(t, "example").
		File(
			"join.go",
			"package example\n\nfunc Join(\n\tmsg string,\n) string {\n\treturn msg\n}\n",
		).
		Directory()
}
