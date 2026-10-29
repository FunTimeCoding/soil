package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
)

func (c *Client) DeleteEventSimple(v *event.Event) error {
	f := filter.New()
	f.AddClusters(v.Cluster)
	f.AddNamespaces(v.Namespace)
	f.AddNames(v.Name)

	return c.DeleteEvent(f)
}
