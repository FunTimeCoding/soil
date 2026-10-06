package gosublime

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/spf13/cobra"
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
				x.Terminal.Exitf("error: %v\n", e)
			}

			if r.JSON200 == nil {
				x.Terminal.Exitf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
			}

			for _, v := range *r.JSON200 {
				dirty := " "

				if v.IsDirty {
					dirty = "*"
				}

				path := v.FilePath

				if path == "" {
					path = "-"
				}

				console.Format(
					"%4d %s %s  %s\n",
					v.ViewId,
					dirty,
					v.Title,
					path,
				)
			}
		},
	}
}
