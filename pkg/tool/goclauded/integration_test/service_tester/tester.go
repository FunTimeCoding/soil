package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
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
	Harbor            string
}
