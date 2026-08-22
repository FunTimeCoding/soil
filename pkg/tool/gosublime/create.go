package gosublime

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	telemetry "github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	gosublimed "github.com/funtimecoding/soil/pkg/tool/gosublimed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/spf13/cobra"
	"os"
)

func create(x *Context) *cobra.Command {
	var syntax string
	result := &cobra.Command{
		Use:   "create <title> <content>",
		Short: "Create a scratch view",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			body := client.CreateViewJSONRequestBody{
				Title:   arguments[0],
				Content: arguments[1],
			}

			if syntax != "" {
				body.Syntax = &syntax
			}

			r, e := x.Client.CreateViewWithResponse(context.Background(), body)

			if e != nil {
				errors.Printf("error: %v\n", e)
				os.Exit(1)
			}

			if r.JSON200 == nil {
				errors.Printf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
				os.Exit(1)
			}

			x.Telemetry.Record(
				record.NewDomain(
					gosublimed.CreateView,
					telemetry.CommandLine,
					telemetry.User,
					telemetry.Success,
				),
			)
			fmt.Printf("created view %d\n", r.JSON200.ViewId)
		},
	}
	result.Flags().StringVar(&syntax, "syntax", "", "syntax name")

	return result
}
