package mock_client

import "testing"

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{nextIdentifier: 1}
	t.Cleanup(result.Close)

	return result
}
