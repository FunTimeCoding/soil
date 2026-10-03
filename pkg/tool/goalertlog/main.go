package goalertlog

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/goalertlog/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/client"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	c := client.NewEnvironment()
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(alerts(c))
	o.AddCommand(topAlerts(c))
	argument.CobraStamp(o, constant.Identity, version, gitHash, buildDate)
	errors.PanicOnError(o.Execute())
}
