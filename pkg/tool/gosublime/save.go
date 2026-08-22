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
	"strconv"
)

func save(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "save <id> [path]",
		Short: "Save a view, optionally to a new path",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])

			if e != nil {
				errors.Printf("invalid id: %s\n", arguments[0])
				os.Exit(1)
			}

			var body client.SaveViewJSONRequestBody

			if len(arguments) > 1 {
				body.FilePath = &arguments[1]
			}

			r, f := x.Client.SaveViewWithResponse(
				context.Background(),
				identifier,
				body,
			)

			if f != nil {
				errors.Printf("error: %v\n", f)
				os.Exit(1)
			}

			if r.HTTPResponse.StatusCode != 204 {
				errors.Printf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
				os.Exit(1)
			}

			x.Telemetry.Record(
				record.NewDomain(
					gosublimed.SaveView,
					telemetry.CommandLine,
					telemetry.User,
					telemetry.Success,
				),
			)
			fmt.Printf("saved view %d\n", identifier)
		},
	}
}
