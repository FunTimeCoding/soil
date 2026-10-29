package mock_client

import "k8s.io/api/networking/v1"

func (c *Client) NetworkPolicy(
	_ string,
	_ string,
) (*v1.NetworkPolicy, error) {
	return nil, c.Failure
}
