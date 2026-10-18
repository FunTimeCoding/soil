package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func listTags(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tags",
		Short: "List all source type tags",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.GetTagsWithResponse(context.Background())
			errors.PanicOnError(e)
			refusal.PanicOnFailure("list tags", r.HTTPResponse, r.Body)

			for _, v := range *r.JSON200 {
				collection := ""

				if v.Collection != nil {
					collection = *v.Collection
				}

				pathPrefix := ""

				if v.PathPrefix != nil {
					pathPrefix = *v.PathPrefix
				}

				sourceType := ""

				if v.SourceType != nil {
					sourceType = *v.SourceType
				}

				console.Format(
					"%s  %s  %s\n",
					collection,
					pathPrefix,
					sourceType,
				)
			}
		},
	}
}
