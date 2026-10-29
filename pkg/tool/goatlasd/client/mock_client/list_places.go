package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
)

func (c *Client) ListPlaces(context.Context) ([]client.Place, error) {
	return c.PlaceList, c.Failure
}
