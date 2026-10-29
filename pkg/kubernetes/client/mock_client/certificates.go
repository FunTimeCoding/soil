package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate"

func (c *Client) Certificates() ([]certificate.Certificate, error) {
	return nil, c.Failure
}
