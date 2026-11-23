package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/spf13/cobra"
)

func remediation(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "remediation",
		Short: "List what to do on the selected systems so their vulnerabilities at the score are addressed; the score defaults to 9.0",
		Args:  cobra.NoArgs,
		Run: func(
			o *cobra.Command,
			_ []string,
		) {
			s := selection(o)
			t.EmitValue(
				c.ListRemediation(
					context.Background(),
					&client.ListRemediationParams{
						Show:   s.Show,
						Label:  s.Label,
						Search: s.Search,
						Score:  s.Score,
					},
				),
			)
		},
	}
	selectionFlags(result)

	return result
}
