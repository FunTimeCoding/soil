package goband

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goband/constant"
	"github.com/spf13/cobra"
)

func status(x *Context) *cobra.Command {
	var instance string
	result := &cobra.Command{
		Use:   "status",
		Short: "Show provisioning, redirection, consent and power state",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			v := readStatus(x, instance)
			t := x.Terminal
			t.Line(fmt.Sprintf("Instance: %s", v.Instance))
			t.Line(fmt.Sprintf("Host: %s", v.Host))
			t.Line(fmt.Sprintf("Provisioning: %s", v.Provisioning))
			t.Line(
				fmt.Sprintf(
					"Screen redirection: enabled=%v timeout=%dm",
					v.ScreenEnabled,
					v.ScreenTimeout,
				),
			)
			t.Line(fmt.Sprintf("Redirection listener: %v", v.Listener))
			t.Line(fmt.Sprintf("Consent: %s", v.Consent))
			t.Line(fmt.Sprintf("Power state: %s", v.Power))
			t.Line(
				fmt.Sprintf(
					"Available power operations: %s",
					join.CommaSpace(v.Available),
				),
			)
		},
	}
	result.Flags().StringVar(
		&instance,
		constant.InstanceArgument,
		"",
		constant.InstanceUsage,
	)

	return result
}
