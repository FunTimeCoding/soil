package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
)

func (c *Client) MustPods(f *filter.Filter) []*pod.Pod {
	result, e := c.Pods(f)
	errors.PanicOnError(e)

	return result
}
