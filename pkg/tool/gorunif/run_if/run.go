package run_if

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/run_if/option"
	"strings"
)

func Run(o *option.RunIf) {
	r := run.New()
	r.Start(
		"git",
		"diff",
		"--name-only",
		fmt.Sprintf("%s..%s", resolveBase(o.Base), o.Head),
	)
	changed := false

	for _, p := range split.NewLine(strings.TrimSpace(r.OutputString)) {
		if o.Verbose {
			fmt.Printf("Change: %s\n", p)
		}

		var match bool

		if o.Suffix {
			match = strings.HasSuffix(p, o.Pattern)
		} else {
			match = strings.HasPrefix(p, o.Pattern)
		}

		if match {
			changed = true

			if o.Verbose {
				fmt.Printf("Match: %s\n", o.Pattern)
			}

			break
		}
	}

	if changed {
		run.New().Execute("sh", "-c", o.Execute)
	}
}
