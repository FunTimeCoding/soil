package gopostgres

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgres/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	v, e := client.NewClient(
		locator.New(
			environment.Fallback(constant.HostEnvironment, web.Localhost),
		).Port(web.ListenPort).Insecure().String(),
	)
	errors.PanicOnError(e)
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(listInstances(v, t))
	o.AddCommand(query(v, t))
	o.AddCommand(explain(v, t))
	o.AddCommand(listSchemas(v, t))
	o.AddCommand(listTables(v, t))
	o.AddCommand(describeTable(v, t))
	o.AddCommand(listIndexes(v, t))
	o.AddCommand(tableSizes(v, t))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
