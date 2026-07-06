package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/namespace"
)

func (c *Client) Namespaces(f *filter.Filter) []*namespace.Namespace {
	var result []*namespace.Namespace

	for _, l := range c.selectClients(f) {
		for _, n := range list.Namespace(l.client, l.context) {
			if f == nil || f.ContainsNamespace(n.Name) {
				result = append(result, namespace.New(&n, l.cluster))
			}
		}
	}

	return result
}
