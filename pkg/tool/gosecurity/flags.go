package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/spf13/cobra"
)

func flags(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "flags",
		Short: "List decisions, newest first; open ones unless --all",
		Args:  cobra.NoArgs,
		Run: func(
			o *cobra.Command,
			_ []string,
		) {
			all, e := o.Flags().GetBool("all")
			cobra.CheckErr(e)
			t.EmitValue(
				c.ListFlags(
					context.Background(),
					&client.ListFlagsParams{IncludeMitigated: &all},
				),
			)
		},
	}
	result.Flags().Bool("all", false, "also list mitigated flags")

	return result
}
