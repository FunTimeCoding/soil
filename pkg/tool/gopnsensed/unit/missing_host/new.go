package missing_host

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/mock_source"

func New() *Client {
	return &Client{Client: mock_source.New()}
}
