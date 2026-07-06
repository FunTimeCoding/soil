package gomemory

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gomemory/constant"
	memoryConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	var host string
	var port int
	var l *client.Client
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
		PersistentPreRun: func(
			_ *cobra.Command,
			_ []string,
		) {
			base := locator.New(host).Port(port).Insecure().String()
			c, e := client.NewClient(base)
			errors.PanicOnError(e)
			l = c
		},
	}
	o.PersistentFlags().StringVar(
		&host,
		"host",
		environment.Fallback(
			memoryConstant.HostEnvironment,
			web.Localhost,
		),
		"gomemoryd host",
	)
	o.PersistentFlags().IntVar(
		&port,
		"port",
		environment.FallbackInteger(
			memoryConstant.PortEnvironment,
			web.ListenPort,
		),
		"gomemoryd port",
	)
	o.AddCommand(profile(&l))
	errors.PanicOnError(o.Execute())
}
