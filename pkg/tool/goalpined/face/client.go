package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/client"
	"io"
)

type Client interface {
	GetPackages(
		x context.Context,
		p *client.GetPackagesParams,
	) ([]client.Listing, error)
	PostPackage(
		x context.Context,
		version string,
		repository string,
		architecture string,
		file string,
		body io.Reader,
	) error
	Reachable(x context.Context) error
	Address() string
	Close()
}
