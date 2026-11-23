package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	gosecuritydConstant "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/spf13/cobra"
)

func upgrades(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "upgrades",
		Short: "List the upgrades that fix findings on the selected systems, the most hosts times vulnerabilities first",
		Args:  cobra.NoArgs,
		Run: func(
			o *cobra.Command,
			_ []string,
		) {
			s := selection(o)
			p := &client.ListUpgradesParams{
				Show:   s.Show,
				Label:  s.Label,
				Search: s.Search,
				Score:  s.Score,
			}
			release, e := o.Flags().GetString(
				gosecuritydConstant.ReleaseParameter,
			)
			cobra.CheckErr(e)

			if release != "" {
				p.Release = &release
			}

			kernel, f := o.Flags().GetBool(gosecuritydConstant.KernelParameter)
			cobra.CheckErr(f)

			if kernel {
				p.Kernel = &kernel
			}

			t.EmitValue(c.ListUpgrades(context.Background(), p))
		},
	}
	selectionFlags(result)
	result.Flags().String(
		gosecuritydConstant.ReleaseParameter,
		"",
		"only this release, e.g. bookworm",
	)
	result.Flags().Bool(
		gosecuritydConstant.KernelParameter,
		false,
		"the kernel upgrades instead of the package upgrades",
	)

	return result
}
