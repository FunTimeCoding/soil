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

func open(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "open <path>",
		Short: "Open a file in Sublime",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := x.Client.OpenFileWithResponse(
				context.Background(),
				client.OpenFileJSONRequestBody{FilePath: arguments[0]},
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
					gosublimed.OpenFile,
					telemetry.CommandLine,
					telemetry.User,
					telemetry.Success,
				),
			)
			fmt.Printf("opened view %d\n", r.JSON200.ViewId)
		},
	}
}
