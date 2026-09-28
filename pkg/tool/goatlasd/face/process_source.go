package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
)

type ProcessSource interface {
	ListProcessesWithResponse(
		q context.Context,
		editors ...client.RequestEditorFn,
	) (*client.ListProcessesResponse, error)
}
