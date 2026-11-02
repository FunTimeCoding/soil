package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

type Client interface {
	ListInstances(x context.Context) ([]client.Instance, error)
	GetStatus(
		x context.Context,
		p *client.GetStatusParams,
	) (*client.Status, error)
	RequestPower(
		x context.Context,
		r client.PowerRequest,
	) (*client.PowerResult, error)
	EnableRedirection(
		x context.Context,
		r client.RedirectionRequest,
	) (*client.RedirectionResult, error)
	Reachable(x context.Context) error
	Address() string
	Close()
}
