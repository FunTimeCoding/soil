package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/delete_operation"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
)

func (c *Client) DeleteEvent(f *filter.Filter) {
	for _, l := range c.selectClients(f) {
		for _, n := range f.Namespaces {
			for _, a := range f.Names {
				delete_operation.Event(l.client, l.context, n, a)
			}
		}
	}
}
