package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/job"
)

func (c *Client) Jobs(namespace string) ([]*job.Job, error) {
	result, e := list.Job(c.client, c.context, namespace, "")

	if e != nil {
		return nil, e
	}

	return job.NewSlice(result, c.cluster), nil
}
