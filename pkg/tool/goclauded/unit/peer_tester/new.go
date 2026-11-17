package peer_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_client"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	queryMock "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"testing"
)

func New(t *testing.T) *Peers {
	t.Helper()
	c := mock_client.New()
	c.Harbor = t.TempDir()

	return &Peers{
		t:          t,
		Client:     c,
		Query:      queryMock.New(t),
		Notifier:   mock_notifier.New(),
		MemoryPeer: memoryMock.New(t),
		Reporter:   memory.New(),
		Harbor:     c.Harbor,
		Source:     t.TempDir(),
	}
}
