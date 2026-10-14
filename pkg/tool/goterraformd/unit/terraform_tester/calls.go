package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"os"
	"strings"
)

func (o *Tester) Calls() []string {
	o.t.Helper()
	b, e := os.ReadFile(o.log)

	if e != nil {
		o.t.Fatal(e)
	}

	return split.NewLine(strings.TrimSpace(string(b)))
}
