package mock_client

import "gitlab.com/gitlab-org/api/client-go/v2"

func New() *Client {
	return &Client{files: make(map[string]*gitlab.File)}
}
