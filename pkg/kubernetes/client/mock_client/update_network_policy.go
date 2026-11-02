package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"k8s.io/api/networking/v1"
)

func (c *Client) UpdateNetworkPolicy(p *v1.NetworkPolicy) error {
	if c.Failure != nil {
		return c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return e
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	k := key(p.Namespace, p.Name)

	if _, okay := c.policies[k]; !okay {
		return not_found.New("network policy", k)
	}

	c.policies[k] = p.DeepCopy()

	return nil
}
