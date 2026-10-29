package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"testing"
)

type Tester struct {
	Service *service.Service
	Indexer *mock_client.Client
	t       *testing.T
}
