package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/job"
)

func (c *Client) Jobs(namespace string) []*job.Job {
	return job.NewSlice(
		list.Job(c.client, c.context, namespace, ""),
		c.cluster,
	)
}
