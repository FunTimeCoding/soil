package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/stamp/report"
	"os"
)

func (i *Instance) Parse() {
	i.flags.Bool(constant.Version, false, "Show version information and exit")

	if i.flags.Lookup(constant.Notation) == nil {
		i.flags.Bool(constant.Notation, false, constant.NotationUsage)
	}

	i.exitOnParseError(i.ParseArguments(os.Args[1:]))
	v, f := i.flags.GetBool(constant.Version)
	errors.PanicOnError(f)

	if !v {
		return
	}

	s := stamp.New()

	if n, g := i.flags.GetBool(constant.Notation); g == nil && n {
		console.Line(notation.MarshalIndent(report.New(i.identity.Name(), s)))
	} else {
		s.Print()
	}

	os.Exit(0)
}
