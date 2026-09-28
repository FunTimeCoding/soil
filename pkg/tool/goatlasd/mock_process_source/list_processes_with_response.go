package mock_process_source

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"net/http"
)

func (c *Client) ListProcessesWithResponse(
	_ context.Context,
	_ ...client.RequestEditorFn,
) (*client.ListProcessesResponse, error) {
	return &client.ListProcessesResponse{
		HTTPResponse: &http.Response{StatusCode: http.StatusOK},
		JSON200:      &c.processes,
	}, nil
}
