package module_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"
)

func (o *Tester) ReplacingUser(
	library string,
	source string,
) string {
	o.t.Helper()

	return module.New(o.t,"example").
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
