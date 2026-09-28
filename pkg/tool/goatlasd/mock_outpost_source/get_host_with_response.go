package mock_outpost_source

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
	"net/http"
)

func (c *Client) GetHostWithResponse(
	_ context.Context,
	_ ...client.RequestEditorFn,
) (*client.GetHostResponse, error) {
	if c.fail != nil {
		return nil, c.fail
	}

	return &client.GetHostResponse{
		HTTPResponse: &http.Response{StatusCode: http.StatusOK},
		JSON200:      &c.host,
	}, nil
}
