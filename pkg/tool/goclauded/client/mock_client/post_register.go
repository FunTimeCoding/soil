package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostRegister(_ context.Context, _ client.PostRegisterJSONRequestBody) (*client.RegisterResponse, error) {
	return &client.RegisterResponse{}, c.Failure
}
