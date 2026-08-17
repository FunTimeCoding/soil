package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/get"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
)

func (c *Client) Lease(
	namespace string,
	name string,
) (*lease.Lease, error) {
	result, e := get.Lease(c.client, c.context, namespace, name)

	if e != nil || result == nil {
		return nil, e
	}

	return lease.New(result, c.cluster), nil
}
