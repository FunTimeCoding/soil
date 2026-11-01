package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
)

type Client interface {
	GetStatus(x context.Context) (client.StatusResponse, error)
	PostRun(
		x context.Context,
		intent string,
	) error
	Reachable(x context.Context) error
	Close()
}
