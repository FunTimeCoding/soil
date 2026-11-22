package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"testing"
)

func New(t *testing.T, address string) *Client {
	t.Helper()
	result := &Client{
		address:           address,
		hostCommandStatus: constant.StoreStatusSuccess,
	}
	result.failures = failure.New(result.PostTrigger)
	t.Cleanup(result.Close)

	return result
}
