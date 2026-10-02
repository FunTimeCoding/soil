package license

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/go_mod/check/license/option"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/monitor"
	item "github.com/funtimecoding/soil/pkg/monitor/constant"
)

func Check(o *option.License) {
	elements := monitor.OnlyConcerns(collect(o), o.All)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	for _, e := range elements {
		console.Line(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(item.GoLicense.Plural)
	}
}
