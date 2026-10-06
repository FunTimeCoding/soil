package goquery

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	c, e := client.NewClient(
		locator.Environment(
			constant.HostEnvironment,
			constant.PortEnvironment,
			constant.InsecureEnvironment,
		).String(),
		client.WithRequestEditorFn(
			web.BearerEditor(environment.Required(constant.TokenEnvironment)),
		),
	)
	errors.PanicOnError(e)
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(search(c))
	o.AddCommand(list(c))
	o.AddCommand(index(c))
	o.AddCommand(addCollection(c))
	o.AddCommand(deleteCollection(c))
	o.AddCommand(embed(c))
	o.AddCommand(addContext(c))
	o.AddCommand(listContexts(c))
	o.AddCommand(listTags(c))
	o.AddCommand(listMetadata(c))
	o.AddCommand(removeContext(c))
	o.AddCommand(status(c))
	t := terminal.New(s)
	o.AddCommand(oversize(c, t))
	o.AddCommand(chunk(c, t))
	o.AddCommand(rechunk(c))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
