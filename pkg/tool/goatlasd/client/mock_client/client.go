package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"

type Client struct {
	PlacementList []client.Placement
	PlaceList     []client.Place
	SightingList  []client.Sighting
	Failure       error
	Closed        int
}
