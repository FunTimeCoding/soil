package goagent

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	goagent "github.com/funtimecoding/soil/pkg/tool/goagent/constant"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	agent "github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(goagent.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	var host string
	var port int
	o := &cobra.Command{
		Use:     goagent.Identity.Usage(),
		Short:   goagent.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
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
		return client.New(
			host,
			port,
			environment.Required(agent.TokenEnvironment),
		)
	}
	o.AddCommand(submit(newClient))
	o.AddCommand(status(newClient))
	errors.PanicOnError(o.Execute())
}
