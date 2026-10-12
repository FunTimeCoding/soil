package gomodule

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/check/stale"
	"github.com/funtimecoding/soil/pkg/go_mod/check/stale/option"
	"github.com/funtimecoding/soil/pkg/tool/gomodule/constant"
	"github.com/spf13/cobra"
)

func staleCommand() *cobra.Command {
	o := option.New()
	result := &cobra.Command{
		Use:   constant.StaleUsage,
		Short: "List dependencies by the age of their latest release, stalest first",
		Args:  cobra.MaximumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			o.Path = soil.CurrentDirectory

			if len(arguments) > 0 {
				o.Path = arguments[0]
			}

			stale.Check(o)
		},
	}
	f := result.Flags()
	f.BoolVar(&o.Notation, argumentConstant.Notation, false, "JSON output")
	f.BoolVar(
		&o.All,
		argumentConstant.All,
		false,
		"Include dependencies released within the threshold",
	)
	f.BoolVar(&o.Verbose, argumentConstant.Verbose, false, "Verbose output")
	f.BoolVar(
		&o.Indirect,
		argumentConstant.Indirect,
		false,
		"Include indirect dependencies",
	)
	f.IntVar(
		&o.Months,
		argumentConstant.Months,
		constant.DefaultMonths,
		"Months without a release before a dependency counts as stale",
	)

	return result
}
