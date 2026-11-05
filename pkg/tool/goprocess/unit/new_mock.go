package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"testing"
)

func newMock(t *testing.T) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	result.ProcessList = []client.Process{
		{Name: "alfa", Running: true},
		{Name: "bravo"},
	}
	t.Cleanup(result.Close)

	return result
}
