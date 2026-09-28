package mock_lease_source

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"net/http"
)

func (c *Client) ListLeasesWithResponse(
	_ context.Context,
	_ *client.ListLeasesParams,
	_ ...client.RequestEditorFn,
) (*client.ListLeasesResponse, error) {
	return &client.ListLeasesResponse{
		HTTPResponse: &http.Response{StatusCode: http.StatusOK},
		JSON200:      &c.leases,
	}, nil
}
