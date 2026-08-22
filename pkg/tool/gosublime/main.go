package gosublime

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/tool/gosublime/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/spf13/cobra"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	host := environment.Optional(constant.HostEnvironment)

	if host == "" {
		host = constant.DefaultHost
	}

	c, e := client.NewClientWithResponses(locator.New(host).Insecure().String())

	if e != nil {
		errors.Printf("client: %v\n", e)
		os.Exit(1)
	}

	x := &Context{Client: c, Telemetry: telemetry.NewEnvironment()}
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(views(x))
	o.AddCommand(read(x))
	o.AddCommand(create(x))
	o.AddCommand(edit(x))
	o.AddCommand(open(x))
	o.AddCommand(save(x))
	o.AddCommand(closeView(x))
	errors.PanicOnError(o.Execute())
}
