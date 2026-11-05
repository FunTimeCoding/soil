package mock_client

import "github.com/funtimecoding/soil/pkg/openbao/status"

func New() *Client {
	return &Client{
		status:  status.New(true, false, "2.7.1"),
		secrets: map[string]map[string]string{},
	}
}
