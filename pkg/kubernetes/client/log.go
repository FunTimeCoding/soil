package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"github.com/funtimecoding/soil/pkg/system"
	"k8s.io/api/core/v1"
)

func (c *Client) Log(
	namespace string,
	pod string,
	container string,
) (string, error) {
	r := c.client.CoreV1().Pods(namespace).GetLogs(
		pod,
		&v1.PodLogOptions{
			Follow:    false,
			Previous:  false,
			Container: container,
		},
	)
	l, e := r.Stream(c.context)

	if e != nil {
		return "", refusal.Classify(e)
	}

	defer errors.LogClose(l)

	return string(system.ReadAll(l)), nil
}
