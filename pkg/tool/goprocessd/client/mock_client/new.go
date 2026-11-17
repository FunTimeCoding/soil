package mock_client

import "testing"

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{Logs: map[string][]string{}, Older: map[string]int{}}
	t.Cleanup(result.Close)

	return result
}
