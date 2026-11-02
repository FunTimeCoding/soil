package double

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_notifier"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	queryMock "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"testing"
)

func New(t *testing.T) *Set {
	t.Helper()
	c := mock_client.New()
	c.Harbor = t.TempDir()

	return &Set{
		t:          t,
		Client:     c,
		Query:      queryMock.New(),
		Notifier:   mock_notifier.New(),
		MemoryPeer: memoryMock.New(),
		Reporter:   memory.New(),
		Harbor:     c.Harbor,
		Source:     t.TempDir(),
	}
}
