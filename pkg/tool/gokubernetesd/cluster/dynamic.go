package cluster

import "k8s.io/client-go/dynamic"

func (c *Cluster) Dynamic() dynamic.Interface {
	return c.dynamic
}
