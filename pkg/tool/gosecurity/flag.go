package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	gosecuritydConstant "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/spf13/cobra"
)

func flag(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "flag <subject>",
		Short: "Record a decision: flag a CVE or a remediation item (e.g. upgrade:bookworm:openssl) as a concern",
		Args:  cobra.ExactArgs(1),
		Run: func(
			o *cobra.Command,
			a []string,
		) {
			b := client.FlagRequest{Identifier: a[0]}
			note, e := o.Flags().GetString(gosecuritydConstant.NoteField)
			cobra.CheckErr(e)

			if note != "" {
				b.Note = &note
			}

			t.EmitValue(c.FlagVulnerability(context.Background(), b))
		},
	}
	result.Flags().String(
		gosecuritydConstant.NoteField,
		"",
		"why it is a concern",
	)

	return result
}
