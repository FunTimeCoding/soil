package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
)

func whole(
	o *option.Analyze,
	patterns []string,
) bool {
	return o.Full || resolve.CoversMainModule(patterns)
}
