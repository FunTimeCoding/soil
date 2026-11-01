package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

type Client interface {
	GetDiff(
		x context.Context,
		module string,
		from string,
		to string,
	) (client.DiffResponse, error)
	Reachable(x context.Context) error
	Close()
}
