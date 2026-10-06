package argument

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	libraryErrors "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/stamp/report"
	"github.com/spf13/pflag"
	"os"
)

func (i *Instance) Parse() {
	i.flags.Bool(constant.Version, false, "Show version information and exit")

	if i.flags.Lookup(constant.Notation) == nil {
		i.flags.Bool(constant.Notation, false, constant.NotationUsage)
	}

	e := i.ParseArguments(os.Args[1:])

	if errors.Is(e, pflag.ErrHelp) {
		os.Exit(0)
	}

	v, f := i.flags.GetBool(constant.Version)
	libraryErrors.PanicOnError(f)

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
