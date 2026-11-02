package service_tester

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	queryMock "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"sync"
	"testing"
)

type Tester struct {
	Store          *store_tester.Tester
	t              *testing.T
	Service        *service.Service
	Client         *mock_client.Client
	Query          *queryMock.Client
	Notifier       *mock_notifier.Notifier
	Memory         *memoryMock.Client
	Search         *search_index.Index
	Reporter       *memory.Memory
	Harbor         string
	Source         string
	searchDatabase *sql.DB
	memoryGuard    *peer.Guard
	queryGuard     *peer.Guard
	stop           *stopper.Stopper
	closeOnce      sync.Once
}
