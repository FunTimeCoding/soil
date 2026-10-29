package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/get"
	"k8s.io/api/networking/v1"
)

func (c *Client) NetworkPolicy(
	namespace string,
	name string,
) (*v1.NetworkPolicy, error) {
	return get.NetworkPolicy(c.client, c.context, namespace, name)
}
