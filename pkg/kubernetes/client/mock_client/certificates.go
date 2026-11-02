package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate"

func (c *Client) Certificates() ([]certificate.Certificate, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	var result []certificate.Certificate

	for _, v := range c.certificates {
		result = append(result, *v)
	}

	return result, nil
}
