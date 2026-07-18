package goclaude

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	web "github.com/funtimecoding/soil/pkg/web/constant"
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
	c := command_context.New()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
		PersistentPreRun: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Initialize(host, port)
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
		environment.FallbackInteger(constant.PortEnvironment, web.ListenPort),
		"goclauded port",
	)
	o.AddCommand(sessionBranch(c))
	o.AddCommand(register(c))
	o.AddCommand(check(c))
	o.AddCommand(guard.New())
	o.AddCommand(wait(c))
	o.AddCommand(usage(c))
	errors.PanicOnError(o.Execute())
}
