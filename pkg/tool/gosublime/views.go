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
)

func views(x *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "views",
		Short: "List open views",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := x.Client.GetViewsWithResponse(context.Background())

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
					gosublimed.ListViews,
					telemetry.CommandLine,
					telemetry.User,
					telemetry.Success,
				),
			)

			for _, v := range *r.JSON200 {
				dirty := " "

				if v.IsDirty {
					dirty = "*"
				}

				path := v.FilePath

				if path == "" {
					path = "-"
				}

				fmt.Printf("%4d %s %s  %s\n", v.ViewId, dirty, v.Title, path)
			}
		},
	}
}
