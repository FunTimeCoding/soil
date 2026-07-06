package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
)

func downloadLocator(c *command_context.Context) *cobra.Command {
	var content string
	var filename string
	result := &cobra.Command{
		Use:   "download-locator [node] [storage] [locator]",
		Short: "Download a file from a URL to a storage backend",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			fmt.Println(
				c.Client().DownloadLocator(
					a[0],
					a[1],
					client.DownloadLocatorJSONRequestBody{
						Content:  content,
						Filename: filename,
						Locator:  a[2],
					},
				),
			)
		},
	}
	result.Flags().StringVar(
		&content,
		"content",
		"import",
		"content type: import, iso, or vztmpl",
	)
	result.Flags().StringVar(
		&filename,
		"filename",
		"",
		"filename to save as",
	)
	errors.PanicOnError(result.MarkFlagRequired("filename"))

	return result
}
