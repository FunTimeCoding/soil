package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/client/mock_client"
	"testing"
)

func secrets(t *testing.T) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	t.Cleanup(result.Close)

	return result
}
