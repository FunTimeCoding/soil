package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client/mock_client"
	"testing"
)

func NewLease(t *testing.T) *mock_client.Client {
	t.Helper()

	return mock_client.New(t)
}
