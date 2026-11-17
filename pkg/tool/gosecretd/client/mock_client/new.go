package mock_client

import "testing"

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{Secrets: map[string]map[string]string{}}
	t.Cleanup(result.Close)

	return result
}
