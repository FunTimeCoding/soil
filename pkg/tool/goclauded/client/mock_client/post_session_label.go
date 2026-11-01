package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostSessionLabel(_ context.Context, _ string, _ client.PostSessionLabelJSONRequestBody) (*client.LabelResponse, error) {
	return &client.LabelResponse{}, c.Failure
}
