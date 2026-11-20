package scope_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
)

func (o *Tester) Caller(
	directory string,
	helperModule string,
) {
	o.t.Helper()
	system_tester.New(o.t).Write(
		directory,
		"caller/caller.go",
		fmt.Sprintf(
			"package caller\n\nimport (\n\t\"%s/helper\"\n\t\"testing\"\n)\n\nfunc UseArrangedElsewhere(t *testing.T) {\n\ts := helper.NewServer(t)\n\ts.Ping()\n}\n",
			helperModule,
		),
	)
}
