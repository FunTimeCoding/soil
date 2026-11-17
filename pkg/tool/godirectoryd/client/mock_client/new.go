package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{Base: "dc=example,dc=org", Passwords: map[string]string{}}
	result.failures = failure.New(result.GetUser)
	t.Cleanup(result.Close)

	return result
}
