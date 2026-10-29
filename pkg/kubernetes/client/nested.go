package client

import "k8s.io/client-go/kubernetes"

func (c *Client) Nested() kubernetes.Interface {
	return c.client
}
