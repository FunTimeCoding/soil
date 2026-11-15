package service_tester

import (
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	"testing"
)

func DeviceCatalog(t *testing.T) *mock_client.Client {
	t.Helper()
	result := NewSource(t)
	result.SeedDeviceRole("server")
	result.SeedDeviceType("r640")
	result.SeedSite("home")

	return result
}
