package goband

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goband/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/spf13/cobra"
)

func redirection(x *Context) *cobra.Command {
	var instance string
	var confirmed bool
	result := &cobra.Command{
		Use:   "redirection",
		Short: "Switch on the redirection listener and screen service where they are off",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			v := readStatus(x, instance)
			pending := pendingRedirection(v)

			if len(pending) == 0 {
				x.Terminal.Line(
					fmt.Sprintf("Redirection already enabled on %s", v.Instance),
				)

				return
			}

			if !confirmed {
				x.Terminal.Line(
					fmt.Sprintf(
						"Would enable %s on %s. Pass --yes to execute.",
						join.CommaSpace(pending),
						v.Instance,
					),
				)

				return
			}

			r, e := x.Client.EnableRedirectionWithResponse(
				context.Background(),
				client.EnableRedirectionJSONRequestBody{Instance: &v.Instance},
			)

			if e != nil {
				x.Terminal.Exitf("error: %v\n", e)
			}

			if r.JSON200 == nil {
				x.Terminal.Reject(r.Status(), r.Body)

				return
			}

			x.Terminal.Line(
				fmt.Sprintf(
					"Enabled %s on %s",
					join.CommaSpace(r.JSON200.Enabled),
					r.JSON200.Instance,
				),
			)
		},
	}
	f := result.Flags()
	f.StringVar(
		&instance,
		constant.InstanceArgument,
		"",
		constant.InstanceUsage,
	)
	f.BoolVar(
		&confirmed,
		constant.YesArgument,
		false,
		"execute the writes instead of describing them",
	)

	return result
}
