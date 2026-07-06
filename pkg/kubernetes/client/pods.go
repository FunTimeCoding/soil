package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"k8s.io/api/core/v1"
)

func (c *Client) Pods(f *filter.Filter) []*pod.Pod {
	var result []*pod.Pod

	if f == nil || len(f.Namespaces) == 0 {
		for _, l := range c.selectClients(f) {
			result = append(
				result,
				pod.NewSlice(
					list.Pod(
						l.client,
						l.context,
						v1.NamespaceAll,
						constant.NodeAll,
					),
					l.cluster,
				)...,
			)
		}
	} else {
		for _, l := range c.selectClients(f) {
			for _, n := range f.Namespaces {
				result = append(
					result,
					pod.NewSlice(
						list.Pod(l.client, l.context, n, constant.NodeAll),
						l.cluster,
					)...,
				)
			}
		}
	}

	return result
}
