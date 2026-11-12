package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

type Client interface {
	GetCensus(
		x context.Context,
		p *client.GetCensusParams,
	) (client.Census, error)
	GetDiff(
		x context.Context,
		p *client.GetDiffParams,
	) (client.DiffResponse, error)
	PostIndexPrune(x context.Context) ([]string, error)
	Reachable(x context.Context) error
	Close()
}
