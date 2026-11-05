package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"io"
)

type Client interface {
	face.Reachable
	GetClaims(x context.Context) ([]client.Claim, error)
	PutClaim(
		x context.Context,
		item string,
		owner string,
	) (*client.Claim, error)
	DeleteClaim(
		x context.Context,
		item string,
		owner string,
	) error
	Listen(x context.Context) (io.ReadCloser, error)
	Close()
}
