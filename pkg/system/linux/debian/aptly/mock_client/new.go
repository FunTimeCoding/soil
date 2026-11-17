package mock_client

import "testing"

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{versions: make(map[string][]string)}
	t.Cleanup(result.Close)

	return result
}
