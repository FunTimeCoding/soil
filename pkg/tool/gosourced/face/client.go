package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

type Client interface {
	GetDiff(
		x context.Context,
		p *client.GetDiffParams,
	) (client.DiffResponse, error)
	GetUsage(
		x context.Context,
		p *client.GetUsageParams,
	) (client.Usage, error)
	PostCheck(
		x context.Context,
		b client.CheckRequest,
	) (client.CheckResponse, error)
	PostIndexPrune(x context.Context) ([]string, error)
	Reachable(x context.Context) error
	Close()
}
