package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"testing"
)

func NewNetbox(t *testing.T) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	t.Cleanup(result.Close)

	return result
}
