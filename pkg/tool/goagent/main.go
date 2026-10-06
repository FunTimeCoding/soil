package goagent

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	goagent "github.com/funtimecoding/soil/pkg/tool/goagent/constant"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	agent "github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(goagent.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	var host string
	var port int
	o := &cobra.Command{
		Use:   goagent.Identity.Usage(),
		Short: goagent.Identity.Description(),
	}
	o.PersistentFlags().StringVar(
		&host,
		argumentConstant.Host,
		constant.Localhost,
		"goagentd host",
	)
	o.PersistentFlags().IntVar(
		&port,
		argumentConstant.Port,
		constant.ListenPort,
		"goagentd port",
	)
	newClient := func() *client.Client {
		return client.New(host, port, t.Required(agent.TokenEnvironment))
	}
	o.AddCommand(submit(newClient, t))
	o.AddCommand(status(newClient))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, goagent.Identity)
	errors.PanicOnError(o.Execute())
}
