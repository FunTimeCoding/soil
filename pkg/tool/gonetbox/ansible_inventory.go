package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/spf13/cobra"
)

func ansibleInventory(
	c *client.Client,
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
			t.Emit(c.AnsibleInventory(field))
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
