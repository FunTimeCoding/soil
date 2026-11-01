package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) DeleteSessionByIdentifier(
	_ context.Context,
	_ string,
	_ *client.DeleteSessionByIdentifierParams,
) (*client.DeleteReceiptResponse, error) {
	return &client.DeleteReceiptResponse{}, c.Failure
}
