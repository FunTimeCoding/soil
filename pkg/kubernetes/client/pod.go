package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/get"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
)

func (c *Client) Pod(
	namespace string,
	name string,
) (*pod.Pod, error) {
	result, e := get.Pod(c.client, c.context, namespace, name)

	if e != nil || result == nil {
		return nil, e
	}

	return pod.New(result, c.cluster), nil
}
