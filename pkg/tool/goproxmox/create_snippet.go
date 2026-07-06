package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
)

func createSnippet(c *command_context.Context) *cobra.Command {
	var content string
	result := &cobra.Command{
		Use:   "create-snippet [name]",
		Short: "Create a snippet file",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			fmt.Println(
				c.Client().CreateSnippet(
					client.CreateSnippetJSONRequestBody{
						Name:    a[0],
						Content: content,
					},
				),
			)
		},
	}
	result.Flags().StringVar(
		&content,
		"content",
		"",
		"snippet content",
	)
	errors.PanicOnError(result.MarkFlagRequired("content"))

	return result
}
