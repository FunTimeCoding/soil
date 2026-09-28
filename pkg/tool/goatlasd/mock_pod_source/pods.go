package mock_pod_source

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
)

func (c *Client) Pods(_ *filter.Filter) []*pod.Pod {
	return c.pods
}
