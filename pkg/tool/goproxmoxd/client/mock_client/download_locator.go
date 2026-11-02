package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) DownloadLocator(
	_ context.Context,
	node string,
	_ string,
	_ *client.DownloadLocatorParams,
	_ client.DownloadLocatorJSONRequestBody,
) (*client.StatusResult, error) {
	if e := c.node(node); e != nil {
		return nil, e
	}

	return &client.StatusResult{Status: "downloading"}, nil
}
