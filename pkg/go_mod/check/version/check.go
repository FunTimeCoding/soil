package version

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version/option"
	"github.com/funtimecoding/soil/pkg/monitor"
)

func Check(o *option.Version) {
	elements := monitor.OnlyConcerns(collect(o), o.All)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := go_mod.Format

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}
}
