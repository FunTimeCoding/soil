package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
	"testing"
)

func NewOutpost(
	t *testing.T,
	hostname string,
	hardwareAddresses []string,
) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	t.Cleanup(result.Close)
	result.HostValue = &client.Host{
		Hostname:          hostname,
		Platform:          "linux",
		HardwareAddresses: &hardwareAddresses,
	}

	return result
}
