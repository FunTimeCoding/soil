package module_tester

import "testing"

func Interacting(t *testing.T) string {
	t.Helper()

	return New(t, "example").
		File(
			"interact.go",
			"package example\n\nfunc Interact(dirName string, dirPath string) string {\n\treturn join(dirName, dirPath, \"some-moderately-long-literal-value-here\")\n}\n\nfunc join(a, b, c string) string { return a + b + c }\n",
		).
		Directory()
}
