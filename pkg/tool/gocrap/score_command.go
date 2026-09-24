package gocrap

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/crap"
	crapConstant "github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/option"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gocrap/constant"
	"github.com/spf13/cobra"
)

func scoreCommand() *cobra.Command {
	o := option.New()
	result := &cobra.Command{
		Use:   constant.ScoreUsage,
		Short: "Score every function by complexity against coverage",
		Args:  cobra.ArbitraryArgs,
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			o.Root = root()
			o.Patterns = patterns(arguments)
			system.ExitOnCode(crap.Run(o))
		},
	}
	f := result.Flags()
	f.BoolVar(&o.Notation, argumentConstant.Notation, false, "JSON output")
	f.BoolVar(&o.Verbose, argumentConstant.Verbose, false, "Verbose output")
	f.BoolVar(
		&o.FailAbove,
		argumentConstant.Fail,
		false,
		"Exit 1 when any function scores above the threshold",
	)
	f.Float64Var(
		&o.Threshold,
		argumentConstant.Threshold,
		crapConstant.DefaultThreshold,
		"Score above which a function is marked",
	)
	f.Float64Var(
		&o.Minimum,
		argumentConstant.Minimum,
		0,
		"Hide functions scoring below this",
	)
	f.IntVar(&o.Top, argumentConstant.Top, 0, "Show only the worst N functions")
	f.StringVar(
		&o.Missing,
		argumentConstant.Missing,
		crapConstant.PolicyPessimistic,
		"Functions without coverage: pessimistic, optimistic or skip",
	)
	f.StringVar(
		&o.Profile,
		argumentConstant.Profile,
		"",
		"Existing coverage profile instead of running go test",
	)
	f.StringVar(
		&o.Mutation,
		argumentConstant.Mutation,
		"",
		"Gremlins JSON report; surviving mutants mark coverage untrusted",
	)
	f.StringVar(
		&o.Baseline,
		argumentConstant.Baseline,
		"",
		"Previous --notation output to compare against",
	)
	f.Float64Var(
		&o.Tolerance,
		argumentConstant.Tolerance,
		crapConstant.DefaultTolerance,
		"Score change below this counts as unchanged",
	)
	f.BoolVar(
		&o.FailRegression,
		argumentConstant.FailRegression,
		false,
		"Exit 1 when any function scored worse than the baseline",
	)
	f.BoolVar(
		&o.IgnoreCovered,
		argumentConstant.IgnoreCovered,
		false,
		"Do not fail on regressions of fully covered functions",
	)
	f.BoolVar(
		&o.All,
		argumentConstant.All,
		false,
		"With a baseline, also show unchanged functions",
	)

	return result
}
