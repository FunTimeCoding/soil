package service_tester

import "github.com/funtimecoding/soil/pkg/netbox/mock_client"

func SeedDeviceCatalog(c *mock_client.Client) {
	c.SeedDeviceRole("server")
	c.SeedDeviceType("r640")
	c.SeedSite("home")
}
