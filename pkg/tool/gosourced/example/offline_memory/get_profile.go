package offline_memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetProfile(
	context.Context,
	*client.GetProfileParams,
) (*client.ProfileResponse, error) {
	return nil, nil
}
