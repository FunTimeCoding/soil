package gomemory

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomemory/constant"
	memoryConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	var host string
	var port int
	var l *client.Client
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
		PersistentPreRun: func(
			_ *cobra.Command,
			_ []string,
		) {
			base := locator.New(host).Port(port).Insecure().String()
			c, e := client.NewClient(
				base,
				client.WithRequestEditorFn(
					web.BearerEditor(
						t.Required(memoryConstant.TokenEnvironment),
					),
				),
			)
			errors.PanicOnError(e)
			l = c
		},
	}
	o.PersistentFlags().StringVar(
		&host,
		"host",
		environment.Fallback(
			memoryConstant.HostEnvironment,
			webConstant.Localhost,
		),
		"gomemoryd host",
	)
	o.PersistentFlags().IntVar(
		&port,
		"port",
		environment.FallbackInteger(
			memoryConstant.PortEnvironment,
			webConstant.ListenPort,
		),
		"gomemoryd port",
	)
	o.AddCommand(profile(&l, t))
	o.AddCommand(statistic(&l, t))
	o.AddCommand(relations(&l, t))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
