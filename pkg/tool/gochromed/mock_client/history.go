package mock_client

import "github.com/funtimecoding/soil/pkg/chromium"

func (c *Client) History(_ string) (*chromium.HistoryResult, error) {
	return nil, nil
}
