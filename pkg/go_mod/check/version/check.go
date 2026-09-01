package version

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version/option"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/monitor"
)

func Check(o *option.Version) {
	elements := monitor.OnlyConcerns(collect(o), o.All)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	for _, e := range elements {
		console.Line(e.Format(f))
	}
}
