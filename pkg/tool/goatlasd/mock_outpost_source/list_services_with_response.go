package mock_outpost_source

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
	"net/http"
)

func (c *Client) ListServicesWithResponse(
	_ context.Context,
	_ *client.ListServicesParams,
	_ ...client.RequestEditorFn,
) (*client.ListServicesResponse, error) {
	return &client.ListServicesResponse{
		HTTPResponse: &http.Response{StatusCode: http.StatusOK},
		JSON200:      &c.services,
	}, nil
}
