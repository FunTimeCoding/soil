package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"k8s.io/api/networking/v1"
)

func (c *Client) NetworkPolicy(
	namespace string,
	name string,
) (*v1.NetworkPolicy, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	p, okay := c.policies[key(namespace, name)]

	if !okay {
		return nil, not_found.New("network policy", key(namespace, name))
	}

	return p.DeepCopy(), nil
}
