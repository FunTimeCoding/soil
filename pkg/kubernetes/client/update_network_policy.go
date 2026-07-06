package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/update_operation"
	"k8s.io/api/networking/v1"
)

func (c *Client) UpdateNetworkPolicy(np *v1.NetworkPolicy) {
	update_operation.NetworkPolicy(c.client, c.context, np)
}
