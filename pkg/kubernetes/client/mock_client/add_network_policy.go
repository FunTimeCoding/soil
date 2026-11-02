package mock_client

import "k8s.io/api/networking/v1"

func (c *Client) AddNetworkPolicy(p *v1.NetworkPolicy) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.policies[key(p.Namespace, p.Name)] = p.DeepCopy()
}
