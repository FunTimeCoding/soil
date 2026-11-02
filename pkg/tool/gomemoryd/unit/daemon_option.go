package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"path/filepath"
	"testing"
)

func daemonOption(
	t *testing.T,
	q *mock_client.Client,
) *option.Memory {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "memory.db")
	result.Query = func() face.Client { return q }

	return result
}
