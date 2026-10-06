package scope_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func Caller(
	t *testing.T,
	directory string,
	helperModule string,
) {
	t.Helper()
	testutil.WriteFile(
		t,
		directory,
		"caller/caller.go",
		fmt.Sprintf(
			"package caller\n\nimport (\n\t\"%s/helper\"\n\t\"testing\"\n)\n\nfunc UseArrangedElsewhere(t *testing.T) {\n\ts := helper.NewServer(t)\n\ts.Ping()\n}\n",
			helperModule,
		),
	)
}
