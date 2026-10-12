package goagent

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goagent/constant"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	goagentdConstant "github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	var host string
	var port int
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.PersistentFlags().StringVar(
		&host,
		argumentConstant.Host,
		webConstant.Localhost,
		"goagentd host",
	)
	o.PersistentFlags().IntVar(
		&port,
		argumentConstant.Port,
		webConstant.ListenPort,
		"goagentd port",
	)
	newClient := func() *client.Client {
		return client.New(
			host,
			port,
			t.Required(goagentdConstant.TokenEnvironment),
		)
	}
	o.AddCommand(submit(newClient, t))
	o.AddCommand(status(newClient))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
