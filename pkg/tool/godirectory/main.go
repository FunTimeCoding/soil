package godirectory

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godirectory/constant"
	daemon "github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
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
	c, e := client.NewClientWithResponses(
		locator.Environment(
			daemon.HostEnvironment,
			daemon.PortEnvironment,
			daemon.InsecureEnvironment,
		).String(),
		client.WithRequestEditorFn(
			web.BearerEditor(environment.Required(daemon.TokenEnvironment)),
		),
	)
	errors.PanicOnError(e)
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(user(c))
	o.AddCommand(group(c))
	errors.PanicOnError(o.Execute())
}
