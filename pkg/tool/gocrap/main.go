package gocrap

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gocrap/constant"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(scoreCommand(t))
	o.AddCommand(attributeCommand(t))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
