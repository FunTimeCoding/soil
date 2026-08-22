package gosublime

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	telemetry "github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	gosublimed "github.com/funtimecoding/soil/pkg/tool/gosublimed/constant"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func read(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "read <id>",
		Short: "Read a view's full text",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])

			if e != nil {
				errors.Printf("invalid id: %s\n", arguments[0])
				os.Exit(1)
			}

			r, e := x.Client.GetViewWithResponse(
				context.Background(),
				identifier,
			)

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
					gosublimed.ReadView,
					telemetry.CommandLine,
					telemetry.User,
					telemetry.Success,
				),
			)

			if r.JSON200.Text != nil {
				fmt.Print(*r.JSON200.Text)
			}
		},
	}
}
