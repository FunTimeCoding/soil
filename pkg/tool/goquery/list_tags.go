package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/spf13/cobra"
)

func listTags(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tags",
		Short: "List all source type tags",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			tags, e := c.GetTags(context.Background())
			t.EmitDone(e)

			for _, v := range tags {
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
