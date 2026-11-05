package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"testing"
)

func newMock(t *testing.T) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	result.Resolved = map[string]string{"alfa": "alfa-identifier"}
	t.Cleanup(result.Close)

	return result
}
