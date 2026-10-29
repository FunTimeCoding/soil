package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/get"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/job"
)

func (c *Client) Job(
	namespace string,
	name string,
) (*job.Job, error) {
	result, e := get.Job(c.client, c.context, namespace, name)

	if e != nil {
		return nil, e
	}

	return job.New(result, c.cluster), nil
}
