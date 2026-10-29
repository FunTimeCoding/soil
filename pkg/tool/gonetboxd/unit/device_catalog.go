package unit

import (
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func deviceCatalog(t *testing.T) *mock_client.Client {
	t.Helper()
	result := service_tester.NewSource(t)
	result.SeedDeviceRole("server")
	result.SeedDeviceType("r640")
	result.SeedSite("home")

	return result
}
