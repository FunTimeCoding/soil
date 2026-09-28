package mock_outpost_source

import "github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"

func New(
	hostname string,
	hardwareAddresses []string,
) *Client {
	return &Client{
		host: client.Host{
			Hostname:          hostname,
			Platform:          "linux",
			HardwareAddresses: &hardwareAddresses,
		},
	}
}
