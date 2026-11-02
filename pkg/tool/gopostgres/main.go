package gopostgres

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgres/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/client"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	var v face.Client
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
		PersistentPreRun: func(
			_ *cobra.Command,
			_ []string,
		) {
			v = client.New(
				locator.Environment(
					constant.HostEnvironment,
					constant.PortEnvironment,
					constant.InsecureEnvironment,
				),
				t.Required(constant.TokenEnvironment),
			)
		},
		PersistentPostRun: func(
			_ *cobra.Command,
			_ []string,
		) {
			v.Close()
		},
	}
	o.AddCommand(listInstances(&v, t))
	o.AddCommand(query(&v, t))
	o.AddCommand(explain(&v, t))
	o.AddCommand(listSchemas(&v, t))
	o.AddCommand(listTables(&v, t))
	o.AddCommand(describeTable(&v, t))
	o.AddCommand(listIndexes(&v, t))
	o.AddCommand(tableSizes(&v, t))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
