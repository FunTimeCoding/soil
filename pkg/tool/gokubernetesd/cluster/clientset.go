package cluster

import "k8s.io/client-go/kubernetes"

func (c *Cluster) Clientset() kubernetes.Interface {
	return c.clientset
}
