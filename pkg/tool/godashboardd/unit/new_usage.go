package unit

import (
	"github.com/funtimecoding/soil/pkg/nextcloud/usage/mock_client"
	"testing"
)

func newUsage(t *testing.T) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	t.Cleanup(result.Close)

	return result
}
