package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func AddMachine(
	f *mock_client.Client,
	identifier int32,
	name string,
	labels *[]*client.Label,
) {
	f.VirtualMachineList = append(
		f.VirtualMachineList,
		&client.VirtualMachine{
			Identifier: identifier,
			Name:       name,
			Labels:     labels,
		},
	)
}
