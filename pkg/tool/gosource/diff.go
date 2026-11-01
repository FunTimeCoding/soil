package gosource

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosource/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/face"
	"github.com/spf13/cobra"
)

func diff(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var module string
	var from string
	var to string
	var text bool
	result := &cobra.Command{
		Use:   "diff",
		Short: "Compare two revisions symbol by symbol, one line per change",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.GetDiff(
				context.Background(),
				diffParameters(module, from, to, text),
			)

			if e != nil {
				t.EmitDone(e)

				return
			}

			for _, l := range DiffLines(r) {
				t.Line(l)
			}
		},
	}
	result.Flags().StringVar(
		&module,
		constant.ModuleFlag,
		"",
		"module name from the gosourced inventory (required)",
	)
	result.Flags().StringVar(
		&from,
		constant.FromFlag,
		"",
		"revision to compare from, HEAD when empty",
	)
	result.Flags().StringVar(
		&to,
		constant.ToFlag,
		"",
		"revision to compare to, the staged index when empty",
	)
	result.Flags().BoolVar(
		&text,
		constant.TextFlag,
		false,
		"print each changed symbol's text under its line",
	)
	errors.PanicOnError(result.MarkFlagRequired(constant.ModuleFlag))

	return result
}
