package cluster

import "k8s.io/client-go/rest"

func (c *Cluster) Configuration() *rest.Config {
	return c.configuration
}
