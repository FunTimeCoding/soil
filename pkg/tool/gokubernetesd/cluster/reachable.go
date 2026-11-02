package cluster

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"k8s.io/client-go/discovery"
)

func (c *Cluster) Reachable(x context.Context) error {
	_, e := discovery.ToDiscoveryInterfaceWithContext(
		c.clientset.Discovery(),
	).ServerVersionWithContext(x)

	return refusal.Classify(e)
}
