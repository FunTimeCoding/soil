package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"os"
	"strings"
)

func (o *Tester) Calls() []string {
	o.t.Helper()
	b, e := os.ReadFile(o.log)

	if e != nil {
		o.t.Fatal(e)
	}

	return split.NewLine(
		constant.FixturePlanPath.ReplaceAllString(
			strings.TrimSpace(string(b)),
			"PLAN",
		),
	)
}
