package gosecurity

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/spf13/cobra"
)

func findings(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "findings",
		Short: "List vulnerabilities across systems with their score and systems per fix kind",
		Args:  cobra.NoArgs,
		Run: func(
			o *cobra.Command,
			_ []string,
		) {
			t.EmitValue(
				c.ListFindings(
					context.Background(),
					&client.ListFindingsParams{Score: score(o)},
				),
			)
		},
	}
	scoreFlag(result)

	return result
}
