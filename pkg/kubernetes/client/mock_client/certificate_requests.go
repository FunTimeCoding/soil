package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate_request"

func (c *Client) CertificateRequests() ([]certificate_request.CertificateRequest, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, c.failures.Caller()
}
