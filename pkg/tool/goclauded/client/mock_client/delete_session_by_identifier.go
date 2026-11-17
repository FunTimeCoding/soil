package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) DeleteSessionByIdentifier(
	_ context.Context,
	_ string,
	_ *client.DeleteSessionByIdentifierParams,
) (*client.DeletionResult, error) {
	return &client.DeletionResult{}, c.Failure
}
