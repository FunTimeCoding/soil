package gosalt

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	saltConstant "github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosalt/constant"
	"github.com/spf13/cobra"
)

func treeFlags(
	c *cobra.Command,
	directories *[]string,
	environments *[]string,
) {
	f := c.Flags()
	f.StringSliceVar(
		directories,
		constant.DirectoryArgument,
		[]string{soil.CurrentDirectory},
		constant.DirectoryUsage,
	)
	f.StringSliceVar(
		environments,
		constant.EnvironmentArgument,
		saltConstant.DefaultEnvironments,
		constant.EnvironmentUsage,
	)
}
