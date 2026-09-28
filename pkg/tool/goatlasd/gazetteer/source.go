package gazetteer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

type Source interface {
	ListDevicesWithResponse(
		q context.Context,
		params *client.ListDevicesParams,
		editors ...client.RequestEditorFn,
	) (*client.ListDevicesResponse, error)
	ListVirtualMachinesWithResponse(
		q context.Context,
		editors ...client.RequestEditorFn,
	) (*client.ListVirtualMachinesResponse, error)
	ListPhysicalAddressesWithResponse(
		q context.Context,
		editors ...client.RequestEditorFn,
	) (*client.ListPhysicalAddressesResponse, error)
}
