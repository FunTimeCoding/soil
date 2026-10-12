package mock_client

import "github.com/funtimecoding/soil/pkg/band/response"

func (c *Client) Consent() (*response.Consent, error) {
	return response.NewConsent(), nil
}
