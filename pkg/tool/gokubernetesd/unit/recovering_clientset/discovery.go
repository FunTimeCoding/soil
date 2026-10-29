package recovering_clientset

import "k8s.io/client-go/discovery"

func (c *Clientset) Discovery() discovery.DiscoveryInterfaces {
	return c.discovery
}
