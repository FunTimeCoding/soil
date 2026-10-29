package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func AddDevice(
	f *mock_client.Client,
	identifier int32,
	name string,
	labels *[]*client.Label,
) {
	f.DeviceList = append(
		f.DeviceList,
		&client.Device{Identifier: identifier, Name: name, Labels: labels},
	)
}
