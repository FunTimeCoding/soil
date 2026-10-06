package goclaude

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	var host string
	var port int
	c := command_context.New(terminal.New(s))
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
		PersistentPreRun: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Initialize(
				host,
				port,
				environment.Exists(constant.InsecureEnvironment),
				c.Terminal().Required(constant.TokenEnvironment),
			)
		},
	}
	o.PersistentFlags().StringVar(
		&host,
		"host",
		environment.Fallback(constant.HostEnvironment, web.Localhost),
		"goclauded host",
	)
	o.PersistentFlags().IntVar(
		&port,
		"port",
		environment.FallbackInteger(constant.PortEnvironment, 0),
		"goclauded port",
	)
	o.AddCommand(sessionBranch(c))
	o.AddCommand(register(c))
	o.AddCommand(sessionEnd(c))
	o.AddCommand(turnEnd(c))
	o.AddCommand(status(c))
	o.AddCommand(check(c))
	o.AddCommand(statusLine(c))
	o.AddCommand(guard.New(c.Terminal()))
	o.AddCommand(serveChannel(c, s.Reporter()))
	o.AddCommand(usage(c))
	o.AddCommand(cost(c))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
