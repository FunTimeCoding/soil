package goquery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func listTags(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tags",
		Short: "List all source type tags",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.GetTags(context.Background())
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			var results []client.SourceTypeTag
			errors.PanicOnError(json.NewDecoder(r.Body).Decode(&results))

			for _, v := range results {
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

				fmt.Printf("%s  %s  %s\n", collection, pathPrefix, sourceType)
			}
		},
	}
}
