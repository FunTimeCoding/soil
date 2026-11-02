package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate"

func (c *Client) AddCertificate(
	namespace string,
	name string,
	issuer string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.certificates = append(
		c.certificates,
		certificate.New(namespace, name, issuer),
	)
}
