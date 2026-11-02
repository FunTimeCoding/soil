package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

type Client interface {
	MachineClient
	ContainerClient
	NodeClient
	SnippetClient
	face.Reachable
	Address() string
	DeriveHardwareAddress(
		x context.Context,
		p *client.DeriveHardwareAddressParams,
	) (*client.DerivedAddress, error)
	ListInstances(x context.Context) ([]client.InstanceDetail, error)
	Reconcile(
		x context.Context,
		p *client.ReconcileParams,
	) (*client.ReconcilePlan, error)
	Close()
}
