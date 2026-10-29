package unit

import (
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"testing"
)

func newPrometheus(t *testing.T) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	t.Cleanup(result.Close)

	return result
}
