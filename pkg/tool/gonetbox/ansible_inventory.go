package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func ansibleInventory(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var field string
	result := &cobra.Command{
		Use:   "ansible-inventory",
		Short: "Print the Ansible inventory NetBox describes, one group per field value",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			markup, e := c.AnsibleInventoryMarkup(context.Background(), field)

			if e != nil {
				t.Exitln(e)

				return
			}

			t.Line(markup)
		},
	}
	result.Flags().StringVar(
		&field,
		constant.Field,
		constant.AnsibleField,
		"custom field whose values become the groups",
	)

	return result
}
