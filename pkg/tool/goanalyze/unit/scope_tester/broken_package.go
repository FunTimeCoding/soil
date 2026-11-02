package scope_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"path/filepath"
	"testing"
)

func BrokenPackage(
	t *testing.T,
	directory string,
	name string,
) {
	t.Helper()
	testutil.WriteFile(
		t,
		directory,
		filepath.Join(name, "broken.go"),
		fmt.Sprintf(
			"package %s\n\nfunc Broken() int {\n\treturn \"alfa\"\n}\n",
			name,
		),
	)
}
