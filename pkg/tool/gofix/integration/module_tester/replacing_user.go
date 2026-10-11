package module_tester

import (
	"fmt"
	"testing"
)

func ReplacingUser(
	t *testing.T,
	library string,
	source string,
) string {
	t.Helper()

	return New(t, "example").
		File(
			"go.mod",
			fmt.Sprintf(
				"module example\n\ngo 1.22\n\nrequire other.test/lib v0.0.0\n\nreplace other.test/lib => %s\n",
				library,
			),
		).
		File("user/use.go", source).
		Directory()
}
