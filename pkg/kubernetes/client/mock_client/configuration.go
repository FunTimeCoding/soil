package mock_client

import "k8s.io/client-go/rest"

func (c *Client) Configuration() *rest.Config {
	return &rest.Config{}
}
