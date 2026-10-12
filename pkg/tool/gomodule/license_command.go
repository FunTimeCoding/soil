package gomodule

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/check/license"
	"github.com/funtimecoding/soil/pkg/go_mod/check/license/option"
	"github.com/funtimecoding/soil/pkg/tool/gomodule/constant"
	"github.com/spf13/cobra"
)

func licenseCommand() *cobra.Command {
	o := option.New()
	result := &cobra.Command{
		Use:   constant.LicenseUsage,
		Short: "Check dependency licenses",
		Args:  cobra.MaximumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			o.Path = soil.CurrentDirectory

			if len(arguments) > 0 {
				o.Path = arguments[0]
			}

			license.Check(o)
		},
	}
	f := result.Flags()
	f.BoolVar(&o.Notation, argumentConstant.Notation, false, "JSON output")
	f.BoolVar(&o.All, argumentConstant.All, false, "Include filtered in output")
	f.BoolVar(&o.Verbose, argumentConstant.Verbose, false, "Verbose output")

	return result
}
