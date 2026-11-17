package refusing_source

import (
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	"testing"
)

func New(t *testing.T, refusal error) *Client {
	t.Helper()
	result := &Client{Client: mock_client.New(t), refusal: refusal}
	t.Cleanup(result.Close)

	return result
}
