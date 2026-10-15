package gobuild

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/constant"
	gobuildConstant "github.com/funtimecoding/soil/pkg/tool/gobuild/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobuild/option"
	"github.com/spf13/cobra"
)

func buildCommand() *cobra.Command {
	o := option.New()
	targets := map[string]*bool{}
	result := &cobra.Command{
		Use:   join.Space(gobuildConstant.BuildCommand, "[name...]"),
		Short: "Build every cmd/ program, or the named ones, for each target",
		RunE: func(
			c *cobra.Command,
			arguments []string,
		) error {
			o.Names = arguments
			o.Targets = selectedTargets(targets)

			if e := Validate(o); e != nil {
				return e
			}

			c.SilenceUsage = true

			return Run(o)
		},
	}
	moduleFlags(result, o)
	f := result.Flags()

	for _, t := range constant.SystemArchitectures {
		targets[t] = f.Bool(t, false, join.Space("Build", t))
	}

	f.StringVar(
		&o.Main,
		gobuildConstant.MainFlag,
		"",
		"Main file, with exactly one name",
	)
	f.StringVar(
		&o.Output,
		gobuildConstant.OutputFlag,
		"",
		"Output path, with exactly one name and one target",
	)

	return result
}
