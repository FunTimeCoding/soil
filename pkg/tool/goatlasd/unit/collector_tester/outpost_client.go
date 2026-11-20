package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
)

func (o *Tester) OutpostClient(
	hostname string,
	hardwareAddresses []string,
) *mock_client.Client {
	o.t.Helper()
	result := mock_client.New(o.t)
	result.HostValue = &client.Host{
		Hostname:          hostname,
		Platform:          "linux",
		HardwareAddresses: &hardwareAddresses,
	}

	return result
}
