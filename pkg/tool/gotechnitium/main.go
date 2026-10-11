package gotechnitium

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gotechnitium/constant"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	c := technitium.NewEnvironment()
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(listZones(c))
	o.AddCommand(createZone(c))
	o.AddCommand(listRecords(c))
	o.AddCommand(addRecord(c))
	o.AddCommand(deleteRecord(c))
	o.AddCommand(listKeys(c))
	o.AddCommand(addKey(c, t))
	o.AddCommand(zoneOptions(c))
	o.AddCommand(allowUpdate(c, t))
	o.AddCommand(allowTransfer(c, t))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
