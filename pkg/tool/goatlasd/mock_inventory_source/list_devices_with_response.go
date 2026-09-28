package mock_inventory_source

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"net/http"
)

func (c *Client) ListDevicesWithResponse(
	_ context.Context,
	_ *client.ListDevicesParams,
	_ ...client.RequestEditorFn,
) (*client.ListDevicesResponse, error) {
	return &client.ListDevicesResponse{
		HTTPResponse: &http.Response{StatusCode: http.StatusOK},
		JSON200:      &c.devices,
	}, nil
}
