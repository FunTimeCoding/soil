package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"k8s.io/api/core/v1"
)

func (c *Client) Pods(f *filter.Filter) ([]*pod.Pod, error) {
	namespaces := []string{v1.NamespaceAll}

	if f != nil && len(f.Namespaces) > 0 {
		namespaces = f.Namespaces
	}

	var result []*pod.Pod

	for _, l := range c.selectClients(f) {
		for _, n := range namespaces {
			v, e := list.Pod(l.client, l.context, n, constant.NodeAll)

			if e != nil {
				return nil, e
			}

			result = append(result, pod.NewSlice(v, l.cluster)...)
		}
	}

	return result, nil
}
