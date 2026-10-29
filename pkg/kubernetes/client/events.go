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
) ([]*event.Event, error) {
	namespaces := []string{v1.NamespaceAll}

	if len(f.Namespaces) > 0 {
		namespaces = f.Namespaces
	}

	var result []*event.Event

	for _, l := range c.selectClients(f) {
		for _, n := range namespaces {
			v, e := list.Event(l.client, l.context, n, limit, fieldSelector)

			if e != nil {
				return nil, e
			}

			result = append(result, event.NewSlice(v, l.cluster)...)
		}
	}

	return event.Sort(result), nil
}
