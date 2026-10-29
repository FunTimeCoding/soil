package mock_client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func (c *Client) Nested() kubernetes.Interface {
	return fake.NewClientset()
}
