package mock_inventory_source

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"net/http"
)

func (c *Client) ListVirtualMachinesWithResponse(
	_ context.Context,
	_ ...client.RequestEditorFn,
) (*client.ListVirtualMachinesResponse, error) {
	return &client.ListVirtualMachinesResponse{
		HTTPResponse: &http.Response{StatusCode: http.StatusOK},
		JSON200:      &c.machines,
	}, nil
}
