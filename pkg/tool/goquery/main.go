package goquery

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	c := client.NewEnvironment()
	defer c.Close()
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(search(c, t))
	o.AddCommand(list(c, t))
	o.AddCommand(index(c, t))
	o.AddCommand(addCollection(c, t))
	o.AddCommand(deleteCollection(c, t))
	o.AddCommand(embed(c, t))
	o.AddCommand(addContext(c, t))
	o.AddCommand(listContexts(c, t))
	o.AddCommand(listTags(c, t))
	o.AddCommand(listMetadata(c, t))
	o.AddCommand(removeContext(c, t))
	o.AddCommand(status(c, t))
	o.AddCommand(oversize(c, t))
	o.AddCommand(chunk(c, t))
	o.AddCommand(rechunk(c, t))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
