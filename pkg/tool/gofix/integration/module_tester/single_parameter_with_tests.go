package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) SingleParameterWithTests() string {
	o.t.Helper()

	return module.New(o.t,"example").
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
