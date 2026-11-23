package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/spf13/cobra"
)

func mitigate(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "mitigate <subject>",
		Short: "Close a decision: a mitigation or an accepted risk with its reason, for a flagged CVE or remediation item",
		Args:  cobra.ExactArgs(1),
		Run: func(
			o *cobra.Command,
			a []string,
		) {
			b := client.MitigationRequest{}
			how, e := o.Flags().GetString("how")
			cobra.CheckErr(e)

			if how != "" {
				b.Mitigation = &how
			}

			t.EmitValue(c.MitigateFlag(context.Background(), a[0], b))
		},
	}
	result.Flags().String("how", "", "how it was mitigated")

	return result
}
