package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_indexer"
	"testing"
)

type Tester struct {
	Store             *store_tester.Tester
	t                 *testing.T
	Service           *service.Service
	Client            *mock_client.Client
	SummaryIndexer    *mock_indexer.Indexer
	CompletionIndexer *mock_indexer.Indexer
	Notifier          *mock_notifier.Notifier
	Memory            *memoryMock.Client
	Search            *search_index.Index
	Reporter          *memory.Memory
	Harbor            string
}
