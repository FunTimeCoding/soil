package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

type LeaseSource interface {
	ListLeasesWithResponse(
		q context.Context,
		params *client.ListLeasesParams,
		editors ...client.RequestEditorFn,
	) (*client.ListLeasesResponse, error)
}
