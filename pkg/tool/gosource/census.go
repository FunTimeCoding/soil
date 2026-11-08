package gosource

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosource/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"github.com/spf13/cobra"
)

func census(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var module string
	var method string
	result := &cobra.Command{
		Use:   "census <package> <type>",
		Short: "Count a struct type's methods - callers, calls, reads, uses and tables, as JSON",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := c.GetCensus(
				context.Background(),
				&client.GetCensusParams{
					Module:  module,
					Package: arguments[0],
					Type:    arguments[1],
				},
			)
			t.EmitValue(CensusMethods(r, method), e)
		},
	}
	result.Flags().StringVar(
		&module,
		constant.ModuleFlag,
		"",
		"module name from the gosourced inventory (required)",
	)
	result.Flags().StringVar(
		&method,
		constant.MethodFlag,
		"",
		"only this method",
	)
	errors.PanicOnError(result.MarkFlagRequired(constant.ModuleFlag))

	return result
}
