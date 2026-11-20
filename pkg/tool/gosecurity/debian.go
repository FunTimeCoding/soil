package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/spf13/cobra"
)

func debian(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "debian <identifier>",
		Short: "Show Debian's fix status for a CVE per source package and release",
		Args:  cobra.ExactArgs(1),
		Run: func(
			o *cobra.Command,
			a []string,
		) {
			p := &client.GetDebianStatusParams{}
			release, e := o.Flags().GetString("release")
			cobra.CheckErr(e)

			if release != "" {
				p.Release = &release
			}

			t.EmitValue(c.GetDebianStatus(context.Background(), a[0], p))
		},
	}
	result.Flags().String("release", "", "only this release, e.g. trixie")

	return result
}
