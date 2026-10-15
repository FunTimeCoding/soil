package gobuild

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gobuild/constant"
	"github.com/spf13/cobra"
)

func Main() {
	r := reporter.NewOptional(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	o := &cobra.Command{
		Use:           constant.Identity.Usage(),
		Short:         constant.Identity.Description(),
		Args:          cobra.NoArgs,
		SilenceErrors: true,
		RunE: func(
			c *cobra.Command,
			_ []string,
		) error {
			return fmt.Errorf("a command is required\n\n%s", c.UsageString())
		},
	}
	o.AddCommand(buildCommand())
	o.AddCommand(installCommand())
	o.AddCommand(staleCommand())
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
