package mock_client

import "k8s.io/api/networking/v1"

func (c *Client) UpdateNetworkPolicy(_ *v1.NetworkPolicy) error {
	return c.Failure
}
