package module_tester

import "testing"

func SingleParameterWithTests(t *testing.T) string {
	t.Helper()

	return New(t, "example").
		File(
			"find.go",
			"package tested\n\nfunc FindLatest(\n\tv []string,\n) *string {\n\treturn nil\n}\n",
		).
		File(
			"find_test.go",
			"package tested\n\nimport \"testing\"\n\nfunc TestFindLatest(t *testing.T) {}\n",
		).
		Directory()
}
