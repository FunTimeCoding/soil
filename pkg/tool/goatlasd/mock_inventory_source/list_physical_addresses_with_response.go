package mock_inventory_source

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"net/http"
)

func (c *Client) ListPhysicalAddressesWithResponse(
	_ context.Context,
	_ ...client.RequestEditorFn,
) (*client.ListPhysicalAddressesResponse, error) {
	return &client.ListPhysicalAddressesResponse{
		HTTPResponse: &http.Response{StatusCode: http.StatusOK},
		JSON200:      &c.physical,
	}, nil
}
