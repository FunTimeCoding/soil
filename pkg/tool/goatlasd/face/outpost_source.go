package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
)

type OutpostSource interface {
	GetHostWithResponse(
		q context.Context,
		editors ...client.RequestEditorFn,
	) (*client.GetHostResponse, error)
	ListServicesWithResponse(
		q context.Context,
		params *client.ListServicesParams,
		editors ...client.RequestEditorFn,
	) (*client.ListServicesResponse, error)
}
