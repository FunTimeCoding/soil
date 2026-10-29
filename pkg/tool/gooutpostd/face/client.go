package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
)

type Client interface {
	face.Reachable
	Address() string
	Services(
		x context.Context,
		origin string,
	) ([]client.Service, error)
	Host(x context.Context) (*client.Host, error)
	Close()
}
