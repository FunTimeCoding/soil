package peer_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	gomemorydMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	goquerydMock "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"path/filepath"
	"testing"
	"time"
)

func NewOption(
	t *testing.T,
) (*option.Option, *gomemorydMock.Client, *goquerydMock.Client) {
	t.Helper()
	d := New(t)
	t.Cleanup(d.Close)

	return d.Option(
		filepath.Join(t.TempDir(), "goclauded.sqlite"),
		func() time.Time { return time.Now().UTC() },
	), d.MemoryPeer, d.Query
}
