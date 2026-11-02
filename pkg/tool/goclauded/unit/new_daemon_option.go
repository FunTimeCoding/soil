package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/double"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	queryMock "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"path/filepath"
	"testing"
	"time"
)

func newDaemonOption(
	t *testing.T,
) (*option.Option, *memoryMock.Client, *queryMock.Client) {
	t.Helper()
	d := double.New(t)
	t.Cleanup(d.Close)

	return d.Option(
		filepath.Join(t.TempDir(), "goclauded.sqlite"),
		func() time.Time { return time.Now().UTC() },
	), d.MemoryPeer, d.Query
}
