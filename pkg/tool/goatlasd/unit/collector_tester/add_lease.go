package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func AddLease(
	f *mock_client.Client,
	hostname string,
	address string,
	hardwareAddress string,
	reserved bool,
) {
	f.LeaseList = append(
		f.LeaseList,
		client.Lease{
			Hostname:        hostname,
			Address:         address,
			HardwareAddress: hardwareAddress,
			Reserved:        reserved,
		},
	)
}
