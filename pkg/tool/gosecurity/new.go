package gosecurity

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurity/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/spf13/cobra"
)

func New(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	result.AddCommand(lookup(c, t))
	result.AddCommand(debian(c, t))
	result.AddCommand(affected(c, t))
	result.AddCommand(findings(c, t))
	result.AddCommand(systems(c, t))
	result.AddCommand(system(c, t))
	result.AddCommand(upgrades(c, t))
	result.AddCommand(upgrade(c, t))
	result.AddCommand(remediation(c, t))
	result.AddCommand(scores(c, t))
	result.AddCommand(flags(c, t))
	result.AddCommand(flag(c, t))
	result.AddCommand(mitigate(c, t))

	return result
}
