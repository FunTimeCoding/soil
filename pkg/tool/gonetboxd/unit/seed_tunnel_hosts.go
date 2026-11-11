package unit

import (
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func seedTunnelHosts(c *mock_client.Client) {
	playbook := map[string]any{constant.AnsibleField: []string{"echo"}}
	alfa := c.SeedAddressedDevice("alfa", "203.0.113.5/32")
	alfa.Site = "lima"
	alfa.CustomFields = playbook
	c.SeedHostAddress("alfa", "203.0.113.5/32", "dcim.interface", 6)
	c.SeedHostAddress("alfa", "198.51.100.1/24", "dcim.interface", 7)
	bravo := c.SeedAddressedDevice("bravo", "192.0.2.6/24")
	bravo.Site = "kilo"
	bravo.CustomFields = playbook
	c.SeedHostAddress("bravo", "198.51.100.4/24", "dcim.interface", 8)
	charlie := c.SeedVirtualMachine("charlie")
	charlie.Site = "lima"
	charlie.CustomFields = playbook
	c.SeedHostAddress(
		"charlie",
		"198.51.100.9/24",
		"virtualization.vminterface",
		7,
	)
	c.SeedTunnelTermination("dcim.interface", 7)
	c.SeedTunnelTermination("dcim.interface", 8)
	c.SeedTunnelTermination("virtualization.vminterface", 7)
}
