package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate_request"

func (c *Client) CertificateRequests() ([]certificate_request.CertificateRequest, error) {
	return nil, c.Failure
}
