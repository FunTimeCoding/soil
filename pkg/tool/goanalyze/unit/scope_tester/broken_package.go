package scope_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
)

func (o *Tester) BrokenPackage(
	directory string,
	name string,
) {
	o.t.Helper()
	system_tester.New(o.t).Write(
		directory,
		filepath.Join(name, "broken.go"),
		fmt.Sprintf(
			"package %s\n\nfunc Broken() int {\n\treturn \"alfa\"\n}\n",
			name,
		),
	)
}
