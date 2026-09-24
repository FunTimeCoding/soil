package run_if

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/run_if/option"
	"strings"
)

func Matches(
	o *option.RunIf,
	files []string,
) bool {
	for _, p := range files {
		if o.Verbose {
			console.Format("Change: %s\n", p)
		}

		var match bool

		if o.Suffix {
			match = strings.HasSuffix(p, o.Pattern)
		} else {
			match = strings.HasPrefix(p, o.Pattern)
		}

		if match {
			if o.Verbose {
				console.Format("Match: %s\n", o.Pattern)
			}

			return true
		}
	}

	return false
}
