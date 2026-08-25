package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
	"k8s.io/api/core/v1"
)

func (c *Client) Events(
	f *filter.Filter,
	limit int,
	fieldSelector string,
) []*event.Event {
	var result []*event.Event

	for _, l := range c.selectClients(f) {
		if len(f.Namespaces) == 0 {
			result = append(
				result,
				event.NewSlice(
					list.Event(
						l.client,
						l.context,
						v1.NamespaceAll,
						limit,
						fieldSelector,
					),
					l.cluster,
				)...,
			)
		} else {
			for _, n := range f.Namespaces {
				result = append(
					result,
					event.NewSlice(
						list.Event(
							l.client,
							l.context,
							n,
							limit,
							fieldSelector,
						),
						l.cluster,
					)...,
				)
			}
		}
	}

	return event.Sort(result)
}
