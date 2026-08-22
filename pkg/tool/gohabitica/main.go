package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/tool/gohabitica/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	x := &Context{
		Client:    client.NewEnvironment(),
		Telemetry: telemetry.NewEnvironment(),
	}
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(tasks(x))
	o.AddCommand(create(x))
	o.AddCommand(score(x))
	o.AddCommand(tags(x))
	o.AddCommand(statistic(x))
	o.AddCommand(cron(x))
	o.AddCommand(allocate(x))
	o.AddCommand(gear(x))
	o.AddCommand(equip(x))
	errors.PanicOnError(o.Execute())
}
