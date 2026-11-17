package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{HostValue: &client.Host{}}
	t.Cleanup(result.Close)

	return result
}
