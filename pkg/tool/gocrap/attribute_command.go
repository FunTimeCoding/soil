package gocrap

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/option"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gocrap/constant"
	"github.com/spf13/cobra"
)

func attributeCommand() *cobra.Command {
	o := option.NewAttribute()
	result := &cobra.Command{
		Use:   constant.AttributeUsage,
		Short: "Report which tests cover which functions",
		Args:  cobra.ArbitraryArgs,
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			o.Root = root()
			o.Patterns = patterns(arguments)
			system.ExitOnCode(crap.RunAttribute(o))
		},
	}
	f := result.Flags()
	f.BoolVar(&o.Notation, argumentConstant.Notation, false, "JSON output")
	f.BoolVar(
		&o.Single,
		argumentConstant.Single,
		false,
		"Only list functions covered by exactly one test",
	)
	f.BoolVar(
		&o.Load,
		argumentConstant.Load,
		false,
		"Only rank tests by functions they alone cover",
	)

	return result
}
