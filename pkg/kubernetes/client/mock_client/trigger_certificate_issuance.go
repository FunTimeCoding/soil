package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) TriggerCertificateIssuance(
	namespace string,
	name string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return e
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, v := range c.certificates {
		if v.Namespace != namespace || v.Name != name {
			continue
		}

		c.triggered = append(c.triggered, key(namespace, name))
		v.Issuing = !c.issueOnTrigger
		v.Ready = c.issueOnTrigger

		return nil
	}

	return not_found.New("certificate", key(namespace, name))
}
