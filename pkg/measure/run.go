package measure

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/measure/option"
	"github.com/funtimecoding/soil/pkg/measure/registry"
	"github.com/funtimecoding/soil/pkg/notation"
)

func Run(o *option.Measure) {
	r := Scan(o, registry.NewDefault().Filter(o.Languages))

	if o.Notation {
		console.Line(notation.MarshalIndent(r))

		return
	}

	if o.ByFile {
		printFiles(r)
	} else {
		printLanguages(r)
	}

	if o.Verbose {
		printUnplaced(r)
	}
}
