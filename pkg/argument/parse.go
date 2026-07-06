package argument

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/build"
	libraryErrors "github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/pflag"
	"os"
)

func (i *Instance) Parse(
	version string,
	gitHash string,
	buildDate string,
) {
	i.flags.Bool(
		Version,
		false,
		"Show version information and exit",
	)
	e := i.flags.Parse(os.Args[1:])

	if errors.Is(e, pflag.ErrHelp) {
		os.Exit(0)
	}

	v, f := i.flags.GetBool(Version)
	libraryErrors.PanicOnError(f)

	if v {
		build.New(version, gitHash, buildDate).Print()
		os.Exit(0)
	}
}
