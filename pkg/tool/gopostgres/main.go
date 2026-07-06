package gopostgres

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gopostgres/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
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
	v, e := client.NewClient(
		locator.New(
			environment.Fallback(
				constant.HostEnvironment,
				web.Localhost,
			),
		).Port(web.ListenPort).Insecure().String(),
	)
	errors.PanicOnError(e)
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(listInstances(v))
	o.AddCommand(query(v))
	o.AddCommand(explain(v))
	o.AddCommand(listSchemas(v))
	o.AddCommand(listTables(v))
	o.AddCommand(describeTable(v))
	o.AddCommand(listIndexes(v))
	o.AddCommand(tableSizes(v))
	errors.PanicOnError(o.Execute())
}
