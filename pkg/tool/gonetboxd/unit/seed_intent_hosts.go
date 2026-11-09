package unit

import (
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func seedIntentHosts(c *mock_client.Client) {
	playbook := map[string]any{constant.AnsibleField: []string{"echo"}}
	alfa := c.SeedVirtualMachine("alfa")
	alfa.Status = constant.StatusActive
	alfa.CustomFields = playbook
	bravo := c.SeedVirtualMachine("bravo")
	bravo.Status = constant.StatusOffline
	bravo.CustomFields = playbook
	charlie := c.SeedDevice("charlie")
	charlie.Status = constant.StatusActive
	charlie.CustomFields = playbook
	delta := c.SeedDevice("delta")
	delta.Status = constant.StatusOffline
	delta.CustomFields = playbook
}
