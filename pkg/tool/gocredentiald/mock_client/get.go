package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/entry_detail"

func (c *Client) Get(_ string) *entry_detail.Detail {
	return nil
}
