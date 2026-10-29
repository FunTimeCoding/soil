package refusing_source

import "github.com/funtimecoding/soil/pkg/netbox/mock_client"

func New(refusal error) *Client {
	return &Client{Client: mock_client.New(), refusal: refusal}
}
