package goprocess

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/goprocess/constant"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	o := &cobra.Command{
		Use:           constant.Identity.Usage(),
		Short:         constant.Identity.Description(),
		SilenceErrors: true,
	}
	o.PersistentFlags().StringP("file", "f", "Procfile", "path to Procfile")
	o.AddCommand(checkCommand())
	o.AddCommand(runCommand())
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
