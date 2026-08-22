package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func deleteChecklistItem(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-checklist-item [key] [index]",
		Short: "Delete a checklist item by one-based index",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			index, e := strconv.Atoi(arguments[1])

			if e != nil {
				errors.Printf("invalid index: %s\n", arguments[1])
				os.Exit(1)
			}

			fmt.Println(c.DeleteChecklistItem(arguments[0], index))
		},
	}
}
