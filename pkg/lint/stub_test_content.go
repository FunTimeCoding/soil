package lint

import "fmt"

func stubTestContent(
	packageName string,
	testdata bool,
	main bool,
	useAssert bool,
) string {
	testName := stubTestName(packageName)

	if main {
		testName = "Stub"
	}

	if testdata || !useAssert {
		return fmt.Sprintf(
			"package %s\n\nimport \"testing\"\n\nfunc Test%s(t *testing.T) {\n\tt.Helper()\n}\n",
			packageName,
			testName,
		)
	}

	return fmt.Sprintf(
		"package %s\n\nimport (\n\t\"github.com/funtimecoding/soil/pkg/assert\"\n\t\"testing\"\n)\n\nfunc Test%s(t *testing.T) {\n\tassert.Stub(t)\n}\n",
		packageName,
		testName,
	)
}
