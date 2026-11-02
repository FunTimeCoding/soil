package argument

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/spf13/pflag"
	"os"
)

func (i *Instance) exitOnParseError(e error) {
	if e == nil {
		return
	}

	if errors.Is(e, pflag.ErrHelp) {
		os.Exit(0)
	}

	writer.Print(os.Stderr, "%s\n\n", e)
	i.PrintUsage()
	os.Exit(2)
}
