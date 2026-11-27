package mock_client

import "testing"

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{}
	t.Cleanup(result.Close)

	return result
}
