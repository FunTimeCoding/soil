package double

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_notifier"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	queryMock "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"testing"
)

type Set struct {
	t          *testing.T
	Client     *mock_client.Client
	Query      *queryMock.Client
	Notifier   *mock_notifier.Notifier
	MemoryPeer *memoryMock.Client
	Reporter   *memory.Memory
	Harbor     string
	Source     string
}
