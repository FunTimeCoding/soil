package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
)

type Client interface {
	face.Reachable
	Address() string
	ListPlacements(
		x context.Context,
		place string,
		source string,
		packageName string,
	) ([]client.Placement, error)
	ListPlaces(x context.Context) ([]client.Place, error)
	ListSightings(
		x context.Context,
		place string,
		unclaimed bool,
	) ([]client.Sighting, error)
	Close()
}
