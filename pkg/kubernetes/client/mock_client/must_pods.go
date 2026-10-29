package mock_client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
)

func (c *Client) MustPods(_ *filter.Filter) []*pod.Pod {
	return nil
}
